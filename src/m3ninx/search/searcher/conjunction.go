// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package searcher

import (
	"sort"

	"github.com/m3db/m3/src/m3ninx/index"
	"github.com/m3db/m3/src/m3ninx/postings"
	"github.com/m3db/m3/src/m3ninx/search"
)

type conjunctionSearcher struct {
	searchers search.Searchers
	negations search.Searchers
}

// NewConjunctionSearcher returns a new Searcher which matches documents which match each
// of the given searchers and none of the negations.
func NewConjunctionSearcher(searchers, negations search.Searchers) (search.Searcher, error) {
	if len(searchers) == 0 {
		return nil, errEmptySearchers
	}

	return &conjunctionSearcher{
		searchers: searchers,
		negations: negations,
	}, nil
}

func (s *conjunctionSearcher) Search(r index.Reader) (postings.List, error) {
	var (
		pl           postings.List
		plNeedsClone = true
	)

	listCount := len(s.searchers)
	if listCount < len(s.negations) {
		listCount = len(s.negations)
	}
	lists := make([]postingsListWithLength, 0, listCount)

	for _, sr := range s.searchers {
		curr, err := sr.Search(r)
		if err != nil {
			return nil, err
		}
		lists = append(lists, postingsListWithLength{
			list:   curr,
			length: curr.Len(),
		})
	}

	sort.Sort(byLengthAscending(lists))
	for _, curr := range lists {
		if pl == nil {
			pl = curr.list
		} else {
			var err error
			pl, err = pl.Intersect(curr.list)
			if err != nil {
				return nil, err
			}
			plNeedsClone = false
		}

		// We can break early if the intersected postings list is ever empty.
		if pl.IsEmpty() {
			break
		}
	}

	lists = lists[:0]
	for _, sr := range s.negations {
		curr, err := sr.Search(r)
		if err != nil {
			return nil, err
		}
		lists = append(lists, postingsListWithLength{
			list:   curr,
			length: curr.Len(),
		})
	}

	sort.Sort(byLengthDescending(lists))
	for _, curr := range lists {
		// We can break early if the resulting postings list is ever empty.
		if pl.IsEmpty() {
			break
		}

		var err error
		pl, err = pl.Difference(curr.list)
		if err != nil {
			return nil, err
		}
		plNeedsClone = false
	}

	if pl != nil && plNeedsClone {
		// There was no new instance created indirectly (by Intersect/Difference), so need to clone.
		pl = pl.CloneAsMutable()
	}

	return pl, nil
}

type postingsListWithLength struct {
	list   postings.List
	length int
}

type byLengthAscending []postingsListWithLength

func (l byLengthAscending) Len() int {
	return len(l)
}

func (l byLengthAscending) Less(i, j int) bool {
	return l[i].length < l[j].length
}

func (l byLengthAscending) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byLengthDescending []postingsListWithLength

func (l byLengthDescending) Len() int {
	return len(l)
}

func (l byLengthDescending) Less(i, j int) bool {
	return l[i].length > l[j].length
}

func (l byLengthDescending) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
