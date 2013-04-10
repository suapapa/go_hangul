package cp949

import (
	"bytes"
	"encoding/binary"
	"sort"
	"sync"
	"unicode/utf8"
)

var (
	fromLookupMutex sync.Mutex
	fromLookupTable cp949Table

	toLookupMutex sync.Mutex
	toLookupTable cp949Table
)

// code pair for a Korean chracter
type cp949Code struct {
	native  uint16 // cp949
	unicode rune   // ucs4
}

// lookup table for translator
type cp949Table []cp949Code

func (t cp949Table) Len() int {
	return len(t)
}

func (t cp949Table) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// instance type to sort the lookup table by native code for from-translator
type cp949TableSortByNative struct{ cp949Table }

func (t cp949TableSortByNative) Less(i, j int) bool {
	return t.cp949Table[i].native < t.cp949Table[j].native
}

// instance type to sort the lookup table by unicode for to-translator
type cp949TableSortByUnicode struct{ cp949Table }

func (t cp949TableSortByUnicode) Less(i, j int) bool {
	return t.cp949Table[i].unicode < t.cp949Table[j].unicode
}

type translator interface {
	Translate(data []byte) (int, []byte, error)
}

// use same struct to from-translator and to-translator.
// each translators use sort.Search() to find corresponding code.
// And, the lookup table is sorted by a filed which have
// same encoding of input data.
type translateCp949 struct {
	table   cp949Table // lookup table
	scratch []byte     // buffer for output
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
			if n <= p.table[i].native {
				return true
			}
			return false
		})

		f := p.table[fi]
		if n == f.native {
			p.scratch = appendRune(p.scratch, f.unicode)
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
		fi := sort.Search(len(p.table), func(i int) bool {
			if r <= p.table[i].unicode {
				return true
			}
			return false
		})

		f := p.table[fi]
		if r == f.unicode {
			p.scratch = append(p.scratch,
				byte(f.native>>8), byte(f.native&0xff))
		} else {
			p.scratch = append(p.scratch, '?')
		}

		data = data[s:]
		c += s
	}
	return c, p.scratch, nil
}

// load cp949.dat to cp949Table
func loadCp949Table() (cp949Table, error) {
	buf := bytes.NewBufferString(cp949LookupData)

	// read info header
	var datInfo struct {
		CodeCnt, ChunkCnt uint16
	}
	if err := binary.Read(buf, binary.BigEndian, &datInfo); err != nil {
		return nil, err
	}

	// read code chunks to table
	table := make(cp949Table, datInfo.CodeCnt)
	table = table[:0]
	var chunk struct {
		Code, Len uint16
	}
	for i := uint16(0); i < datInfo.ChunkCnt; i++ {
		if err := binary.Read(buf, binary.BigEndian, &chunk); err != nil {
			return nil, err
		}

		line := make([]byte, chunk.Len)
		if n, err := buf.Read(line); n != int(chunk.Len) || err != nil {
			return nil, err
		}

		for _, u := range string(line) {
			table = append(table,
				cp949Code{native: chunk.Code, unicode: u})
			chunk.Code += 1
		}
	}

	return table, nil
}

func loadFromCp949LookupTable() (cp949Table, error) {
	fromLookupMutex.Lock()
	defer fromLookupMutex.Unlock()

	if fromLookupTable != nil {
		return fromLookupTable, nil
	}

	t, err := loadCp949Table()
	if err != nil {
		return nil, err
	}
	if !sort.IsSorted(cp949TableSortByNative{t}) {
		panic("cp949.dat is not sorted by native code!")
	}

	fromLookupTable = t
	return fromLookupTable, nil
}

func loadToCp949LookupTable() (cp949Table, error) {
	toLookupMutex.Lock()
	defer toLookupMutex.Unlock()

	if toLookupTable != nil {
		return toLookupTable, nil
	}

	var t cp949Table
	copy(t, fromLookupTable)
	sort.Sort(cp949TableSortByUnicode{t})

	toLookupTable = t
	return toLookupTable, nil
}

func fromCp949() (translator, error) {
	t, err := loadFromCp949LookupTable()
	if err != nil {
		return nil, err
	}
	return &translateFromCp949{table: t}, nil
}

func toCp949() (translator, error) {
	t, err := loadToCp949LookupTable()
	if err != nil {
		return nil, err
	}
	return &translateToCp949{table: t}, nil
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
