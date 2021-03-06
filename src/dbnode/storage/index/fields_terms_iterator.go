// Copyright (c) 2019 Uber Technologies, Inc.
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

package index

import (
	"errors"

	pilosaroaring "github.com/m3dbx/pilosa/roaring"

	"github.com/m3db/m3/src/dbnode/tracepoint"
	"github.com/m3db/m3/src/m3ninx/index/segment"
	"github.com/m3db/m3/src/m3ninx/postings"
	"github.com/m3db/m3/src/m3ninx/postings/roaring"
	"github.com/m3db/m3/src/m3ninx/search"
	"github.com/m3db/m3/src/x/context"
	xerrors "github.com/m3db/m3/src/x/errors"
)

var errUnpackBitmapFromPostingsList = errors.New("unable to unpack bitmap from postings list")

// fieldsAndTermsIteratorOpts configures the fieldsAndTermsIterator.
type fieldsAndTermsIteratorOpts struct {
	restrictByQuery *Query
	iterateTerms    bool
	allowFn         allowFn
	fieldIterFn     newFieldIterFn
}

func (o fieldsAndTermsIteratorOpts) allow(f []byte) bool {
	if o.allowFn == nil {
		return true
	}
	return o.allowFn(f)
}

func (o fieldsAndTermsIteratorOpts) newFieldIter(r segment.Reader) (segment.FieldsPostingsListIterator, error) {
	if o.fieldIterFn == nil {
		return r.FieldsPostingsList()
	}
	return o.fieldIterFn(r)
}

type allowFn func(field []byte) bool

type newFieldIterFn func(r segment.Reader) (segment.FieldsPostingsListIterator, error)

type fieldsAndTermsIter struct {
	reader segment.Reader
	opts   fieldsAndTermsIteratorOpts

	err       error
	fieldIter segment.FieldsPostingsListIterator
	termIter  segment.TermsIterator

	current struct {
		field    []byte
		term     []byte
		postings postings.List
	}

	restrictByPostings *pilosaroaring.Bitmap
}

var fieldsAndTermsIterZeroed fieldsAndTermsIter

var _ fieldsAndTermsIterator = &fieldsAndTermsIter{}

// newFieldsAndTermsIteratorFn is the lambda definition of the ctor for fieldsAndTermsIterator.
type newFieldsAndTermsIteratorFn func(
	ctx context.Context, r segment.Reader, opts fieldsAndTermsIteratorOpts,
) (fieldsAndTermsIterator, error)

func newFieldsAndTermsIterator(
	ctx context.Context,
	reader segment.Reader,
	opts fieldsAndTermsIteratorOpts,
) (fieldsAndTermsIterator, error) {
	iter := &fieldsAndTermsIter{
		reader: reader,
		opts:   opts,
	}

	fiter, err := iter.opts.newFieldIter(reader)
	if err != nil {
		return nil, err
	}
	iter.fieldIter = fiter

	if opts.restrictByQuery == nil {
		// No need to restrict results by query.
		return iter, nil
	}

	// If need to restrict by query, run the query on the segment first.
	searchQuery := opts.restrictByQuery.SearchQuery()
	searcher, err := searchQuery.Searcher()
	if err != nil {
		return nil, err
	}

	var (
		_, sp = ctx.StartTraceSpan(tracepoint.FieldTermsIteratorIndexSearch)
		pl    postings.List
	)
	if readThrough, ok := reader.(search.ReadThroughSegmentSearcher); ok {
		pl, err = readThrough.Search(searchQuery, searcher)
	} else {
		pl, err = searcher.Search(reader)
	}
	sp.Finish()
	if err != nil {
		return nil, err
	}

	// Hold onto the postings bitmap to intersect against on a per term basis.
	bitmap, ok := roaring.BitmapFromPostingsList(pl)
	if !ok {
		return nil, errUnpackBitmapFromPostingsList
	}

	iter.restrictByPostings = bitmap
	return iter, nil
}

func (fti *fieldsAndTermsIter) setNextField() bool {
	fieldIter := fti.fieldIter
	if fieldIter == nil {
		return false
	}

	for fieldIter.Next() {
		field, pl := fieldIter.Current()
		if !fti.opts.allow(field) {
			continue
		}
		if fti.restrictByPostings == nil {
			// No restrictions.
			fti.current.field = field
			return true
		}

		bitmap, ok := roaring.BitmapFromPostingsList(pl)
		if !ok {
			fti.err = errUnpackBitmapFromPostingsList
			return false
		}

		// Check field is part of at least some of the documents we're
		// restricted to providing results for based on intersection
		// count.
		// Note: IntersectionCount is significantly faster than intersecting and
		// counting results and also does not allocate.
		if n := fti.restrictByPostings.IntersectionCount(bitmap); n < 1 {
			// No match, not next result.
			continue
		}

		// Matches, this is next result.
		fti.current.field = field
		return true
	}

	fti.err = fieldIter.Err()
	return false
}

func (fti *fieldsAndTermsIter) setNext() bool {
	// check if current field has another term
	if fti.termIter != nil {
		hasNextTerm, err := fti.nextTermsIterResult()
		if err != nil {
			fti.err = err
			return false
		}
		if hasNextTerm {
			return true
		}
	}

	// i.e. need to switch to next field
	for hasNextField := fti.setNextField(); hasNextField; hasNextField = fti.setNextField() {
		// and get next term for the field
		var err error
		fti.termIter, err = fti.reader.Terms(fti.current.field)
		if err != nil {
			fti.err = err
			return false
		}

		hasNextTerm, err := fti.nextTermsIterResult()
		if err != nil {
			fti.err = err
			return false
		}
		if hasNextTerm {
			return true
		}
	}

	// Check field iterator did not encounter error.
	if err := fti.fieldIter.Err(); err != nil {
		fti.err = err
		return false
	}

	// No more fields.
	return false
}

func (fti *fieldsAndTermsIter) nextTermsIterResult() (bool, error) {
	for fti.termIter.Next() {
		fti.current.term, fti.current.postings = fti.termIter.Current()
		if fti.restrictByPostings == nil {
			// No restrictions.
			return true, nil
		}

		bitmap, ok := roaring.BitmapFromPostingsList(fti.current.postings)
		if !ok {
			return false, errUnpackBitmapFromPostingsList
		}

		// Check term is part of at least some of the documents we're
		// restricted to providing results for based on intersection
		// count.
		// Note: IntersectionCount is significantly faster than intersecting and
		// counting results and also does not allocate.
		if n := fti.restrictByPostings.IntersectionCount(bitmap); n > 0 {
			// Matches, this is next result.
			return true, nil
		}
	}
	if err := fti.termIter.Err(); err != nil {
		return false, err
	}
	if err := fti.termIter.Close(); err != nil {
		return false, err
	}
	// Term iterator no longer relevant, no next.
	fti.termIter = nil
	return false, nil
}

func (fti *fieldsAndTermsIter) Next() bool {
	if fti.err != nil {
		return false
	}
	// if only need to iterate fields
	if !fti.opts.iterateTerms {
		return fti.setNextField()
	}
	// iterating both fields and terms
	return fti.setNext()
}

func (fti *fieldsAndTermsIter) Current() (field, term []byte) {
	return fti.current.field, fti.current.term
}

func (fti *fieldsAndTermsIter) Err() error {
	return fti.err
}

func (fti *fieldsAndTermsIter) Close() error {
	var multiErr xerrors.MultiError
	if fti.fieldIter != nil {
		multiErr = multiErr.Add(fti.fieldIter.Close())
	}
	if fti.termIter != nil {
		multiErr = multiErr.Add(fti.termIter.Close())
	}
	return multiErr.FinalError()
}
