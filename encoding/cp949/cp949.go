// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp949

import (
	"sort"
	"unicode/utf8"
)

type lookupItem struct {
	cp949, ucs2 uint16
}

type lookupTable []lookupItem

type translator interface {
	Translate(data []byte) (int, []byte, error)
}

// use same struct to from-translator and to-translator.
// each translators use sort.Search() to find corresponding code.
// And, the lookup table is sorted by a filed which have
// same encoding of input data.
type translateCp949 struct {
	table   lookupTable // lookup table
	scratch []byte      // buffer for output
}

// from cp949 to unicode translator
type translateFromCp949 translateCp949

func (p *translateFromCp949) Translate(data []byte) (int, []byte, error) {
	p.scratch = p.scratch[:0]
	c := 0
	for len(data) > 0 {
		if data[0]&0x80 == 0 {
			p.scratch = append(p.scratch, data[0])
			data = data[1:]
			c += 1
			continue
		}

		n := uint16(data[0])<<8 | uint16(data[1])
		fi := sort.Search(len(p.table), func(i int) bool {
			if n <= p.table[i].cp949 {
				return true
			}
			return false
		})

		if fi < len(p.table) && n == p.table[fi].cp949 {
			f := p.table[fi]
			p.scratch = appendRune(p.scratch, rune(f.ucs2))
		} else {
			p.scratch = appendRune(p.scratch, utf8.RuneError)
		}
		data = data[2:]
		c += 2
	}
	return c, p.scratch, nil
}

// from unicode to cp949 translator
type translateToCp949 translateCp949

func (p *translateToCp949) Translate(data []byte) (int, []byte, error) {
	p.scratch = p.scratch[:0]
	c := 0
	for len(data) > 0 {
		if data[0]&0x80 == 0 {
			p.scratch = append(p.scratch, data[0])
			data = data[1:]
			c += 1
			continue
		}

		r, s := utf8.DecodeRune(data)
		rUcs2 := uint16(r)
		fi := sort.Search(len(p.table), func(i int) bool {
			if rUcs2 <= p.table[i].ucs2 {
				return true
			}
			return false
		})

		if fi >= len(p.table) {
			p.scratch = append(p.scratch, '?')
		}

		if fi < len(p.table) && rUcs2 == p.table[fi].ucs2 {
			f := p.table[fi]
			p.scratch = append(p.scratch,
				byte(f.cp949>>8), byte(f.cp949&0xff))
		} else {
			p.scratch = append(p.scratch, '?')
		}

		data = data[s:]
		c += s
	}
	return c, p.scratch, nil
}

func fromCp949() (translator, error) {
	return &translateFromCp949{table: fromTable}, nil
}

func toCp949() (translator, error) {
	return &translateToCp949{table: toTable}, nil
}

// ensureCap returns s with a capacity of at least n bytes.
// If cap(s) < n, then it returns a new copy of s with the
// required capacity.
func ensureCap(s []byte, n int) []byte {
	if n <= cap(s) {
		return s
	}
	// logic adapted from appendslice1 in runtime
	m := cap(s)
	if m == 0 {
		m = n
	} else {
		for {
			if m < 1024 {
				m += m
			} else {
				m += m / 4
			}
			if m >= n {
				break
			}
		}
	}
	t := make([]byte, len(s), m)
	copy(t, s)
	return t
}

// XXX: can it be simpler?
func appendRune(buf []byte, r rune) []byte {
	n := len(buf)
	buf = ensureCap(buf, n+utf8.UTFMax)
	nu := utf8.EncodeRune(buf[n:n+utf8.UTFMax], r)
	return buf[0 : n+nu]
}
