// Copyright (c) 2021 Uber Technologies, Inc.
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

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package cache

// Copyright (c) 2021 Uber Technologies, Inc.
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

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

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

// namespaceResultsMapHash is the hash for a given map entry, this is public to support
// iterating over the map using a native Go for loop.
type namespaceResultsMapHash uint64

// namespaceResultsMapHashFn is the hash function to execute when hashing a key.
type namespaceResultsMapHashFn func([]byte) namespaceResultsMapHash

// namespaceResultsMapEqualsFn is the equals key function to execute when detecting equality of a key.
type namespaceResultsMapEqualsFn func([]byte, []byte) bool

// namespaceResultsMapCopyFn is the copy key function to execute when copying the key.
type namespaceResultsMapCopyFn func([]byte) []byte

// namespaceResultsMapFinalizeFn is the finalize key function to execute when finished with a key.
type namespaceResultsMapFinalizeFn func([]byte)

// namespaceResultsMap uses the genny package to provide a generic hash map that can be specialized
// by running the following command from this root of the repository:
// ```
// make hashmap-gen pkg=outpkg key_type=Type value_type=Type out_dir=/tmp
// ```
// Or if you would like to use bytes or ident.ID as keys you can use the
// partially specialized maps to generate your own maps as well:
// ```
// make byteshashmap-gen pkg=outpkg value_type=Type out_dir=/tmp
// make idhashmap-gen pkg=outpkg value_type=Type out_dir=/tmp
// ```
// This will output to stdout the generated source file to use for your map.
// It uses linear probing by incrementing the number of the hash created when
// hashing the identifier if there is a collision.
// namespaceResultsMap is a value type and not an interface to allow for less painful
// upgrades when adding/removing methods, it is not likely to need mocking so
// an interface would not be super useful either.
type namespaceResultsMap struct {
	_namespaceResultsMapOptions

	// lookup uses hash of the identifier for the key and the MapEntry value
	// wraps the value type and the key (used to ensure lookup is correct
	// when dealing with collisions), we use uint64 for the hash partially
	// because lookups of maps with uint64 keys has a fast path for Go.
	lookup map[namespaceResultsMapHash]namespaceResultsMapEntry
}

// _namespaceResultsMapOptions is a set of options used when creating an identifier map, it is kept
// private so that implementers of the generated map can specify their own options
// that partially fulfill these options.
type _namespaceResultsMapOptions struct {
	// hash is the hash function to execute when hashing a key.
	hash namespaceResultsMapHashFn
	// equals is the equals key function to execute when detecting equality.
	equals namespaceResultsMapEqualsFn
	// copy is the copy key function to execute when copying the key.
	copy namespaceResultsMapCopyFn
	// finalize is the finalize key function to execute when finished with a
	// key, this is optional to specify.
	finalize namespaceResultsMapFinalizeFn
	// initialSize is the initial size for the map, use zero to use Go's std map
	// initial size and consequently is optional to specify.
	initialSize int
}

// namespaceResultsMapEntry is an entry in the map, this is public to support iterating
// over the map using a native Go for loop.
type namespaceResultsMapEntry struct {
	// key is used to check equality on lookups to resolve collisions
	key _namespaceResultsMapKey
	// value type stored
	value results
}

type _namespaceResultsMapKey struct {
	key      []byte
	finalize bool
}

// Key returns the map entry key.
func (e namespaceResultsMapEntry) Key() []byte {
	return e.key.key
}

// Value returns the map entry value.
func (e namespaceResultsMapEntry) Value() results {
	return e.value
}

// _namespaceResultsMapAlloc is a non-exported function so that when generating the source code
// for the map you can supply a public constructor that sets the correct
// hash, equals, copy, finalize options without users of the map needing to
// implement them themselves.
func _namespaceResultsMapAlloc(opts _namespaceResultsMapOptions) *namespaceResultsMap {
	m := &namespaceResultsMap{_namespaceResultsMapOptions: opts}
	m.Reallocate()
	return m
}

func (m *namespaceResultsMap) newMapKey(k []byte, opts _namespaceResultsMapKeyOptions) _namespaceResultsMapKey {
	key := _namespaceResultsMapKey{key: k, finalize: opts.finalizeKey}
	if !opts.copyKey {
		return key
	}

	key.key = m.copy(k)
	return key
}

func (m *namespaceResultsMap) removeMapKey(hash namespaceResultsMapHash, key _namespaceResultsMapKey) {
	delete(m.lookup, hash)
	if key.finalize {
		m.finalize(key.key)
	}
}

// Get returns a value in the map for an identifier if found.
func (m *namespaceResultsMap) Get(k []byte) (results, bool) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			return entry.value, true
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}
	var empty results
	return empty, false
}

// Set will set the value for an identifier.
func (m *namespaceResultsMap) Set(k []byte, v results) {
	m.set(k, v, _namespaceResultsMapKeyOptions{
		copyKey:     true,
		finalizeKey: m.finalize != nil,
	})
}

// namespaceResultsMapSetUnsafeOptions is a set of options to use when setting a value with
// the SetUnsafe method.
type namespaceResultsMapSetUnsafeOptions struct {
	NoCopyKey     bool
	NoFinalizeKey bool
}

// SetUnsafe will set the value for an identifier with unsafe options for how
// the map treats the key.
func (m *namespaceResultsMap) SetUnsafe(k []byte, v results, opts namespaceResultsMapSetUnsafeOptions) {
	m.set(k, v, _namespaceResultsMapKeyOptions{
		copyKey:     !opts.NoCopyKey,
		finalizeKey: !opts.NoFinalizeKey,
	})
}

type _namespaceResultsMapKeyOptions struct {
	copyKey     bool
	finalizeKey bool
}

func (m *namespaceResultsMap) set(k []byte, v results, opts _namespaceResultsMapKeyOptions) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			m.lookup[hash] = namespaceResultsMapEntry{
				key:   entry.key,
				value: v,
			}
			return
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}

	m.lookup[hash] = namespaceResultsMapEntry{
		key:   m.newMapKey(k, opts),
		value: v,
	}
}

// Iter provides the underlying map to allow for using a native Go for loop
// to iterate the map, however callers should only ever read and not write
// the map.
func (m *namespaceResultsMap) Iter() map[namespaceResultsMapHash]namespaceResultsMapEntry {
	return m.lookup
}

// Len returns the number of map entries in the map.
func (m *namespaceResultsMap) Len() int {
	return len(m.lookup)
}

// Contains returns true if value exists for key, false otherwise, it is
// shorthand for a call to Get that doesn't return the value.
func (m *namespaceResultsMap) Contains(k []byte) bool {
	_, ok := m.Get(k)
	return ok
}

// Delete will remove a value set in the map for the specified key.
func (m *namespaceResultsMap) Delete(k []byte) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			m.removeMapKey(hash, entry.key)
			return
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}
}

// Reset will reset the map by simply deleting all keys to avoid
// allocating a new map.
func (m *namespaceResultsMap) Reset() {
	for hash, entry := range m.lookup {
		m.removeMapKey(hash, entry.key)
	}
}

// Reallocate will avoid deleting all keys and reallocate a new
// map, this is useful if you believe you have a large map and
// will not need to grow back to a similar size.
func (m *namespaceResultsMap) Reallocate() {
	if m.initialSize > 0 {
		m.lookup = make(map[namespaceResultsMapHash]namespaceResultsMapEntry, m.initialSize)
	} else {
		m.lookup = make(map[namespaceResultsMapHash]namespaceResultsMapEntry)
	}
}
