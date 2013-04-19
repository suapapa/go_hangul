// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp949

import (
	"bytes"
	"sort"
	"testing"
)

type cp949Test struct {
	cp949, utf8 string
}

var tests = []cp949Test{
	cp949Test{
		cp949: "\xbe\xc6\xb8\xa7\xb4\xd9\xbf\xee \xbf\xec\xb8\xae\xb8\xbb",
		utf8:  "아름다운 우리말",
	},
	cp949Test{
		cp949: "\x8cc\xb9\xe6\xb0\xa2\xc7\xcf",
		utf8:  "똠방각하",
	},
	cp949Test{
		cp949: "\xc6\xe9\xbd\xc3\xc4\xdd\xb6\xf3",
		utf8:  "펩시콜라",
	},
	cp949Test{
		cp949: "\xa8\xc0\xa8\xc0\xb3\xb3!! \xec\xd7\xce\xfa\xea\xc5\xc6\xd0\x92\xe6\x90p\xb1\xc5 \xa8\xde\xa8\xd3\xc4R\xa2\xaf\xa2\xaf\xa2\xaf \xb1\xe0\x8a\x96 \xa8\xd1\xb5\xb3 \xa8\xc0. .\n\xe4\xac\xbf\xb5\xa8\xd1\xb4\xc9\xc8\xc2 . . . . \xbc\xad\xbf\xef\xb7\xef \xb5\xaf\xc7\xd0\xeb\xe0 \xca\xab\xc4R ! ! !\xa4\xd0.\xa4\xd0\n\xc8\xe5\xc8\xe5\xc8\xe5 \xa4\xa1\xa4\xa1\xa4\xa1\xa1\xd9\xa4\xd0_\xa4\xd0 \xbe\xee\x90\x8a \xc5\xcb\xc4\xe2\x83O \xb5\xae\xc0\xc0 \xafh\xce\xfa\xb5\xe9\xeb\xe0 \xa8\xc0\xb5\xe5\x83O\n\xbc\xb3\x90j \xca\xab\xc4R . . . . \xb1\xbc\xbe\xd6\x9af \xa8\xd1\xb1\xc5 \xa8\xde\x90t\xa8\xc2\x83O \xec\xd7\xec\xd2\xf4\xb9\xe5\xfc\xf1\xe9\xb1\xee\xa3\x8e\n\xbf\xcd\xbe\xac\xc4R ! ! \xe4\xac\xbf\xb5\xa8\xd1 \xca\xab\xb4\xc9\xb1\xc5 \xa1\xd9\xdf\xbe\xb0\xfc \xbe\xf8\xb4\xc9\xb1\xc5\xb4\xc9 \xe4\xac\xb4\xc9\xb5\xd8\xc4R \xb1\xdb\xbe\xd6\x8a\xdb\n\xa8\xde\xb7\xc1\xb5\xe0\xce\xfa \x9a\xc3\xc7\xb4\xbd\xa4\xc4R \xbe\xee\x90\x8a \xec\xd7\xec\xd2\xf4\xb9\xe5\xfc\xf1\xe9\x9a\xc4\xa8\xef\xb5\xe9\x9d\xda!! \xa8\xc0\xa8\xc0\xb3\xb3\xa2\xbd \xa1\xd2\xa1\xd2*",
		utf8: `㉯㉯납!! 因九月패믤릔궈 ⓡⓖ훀¿¿¿ 긍뒙 ⓔ뎨 ㉯. .
亞영ⓔ능횹 . . . . 서울뤄 뎐학乙 家훀 ! ! !ㅠ.ㅠ
흐흐흐 ㄱㄱㄱ☆ㅠ_ㅠ 어릨 탸콰긐 뎌응 칑九들乙 ㉯드긐
설릌 家훀 . . . . 굴애쉌 ⓔ궈 ⓡ릘㉱긐 因仁川女中까즼
와쒀훀 ! ! 亞영ⓔ 家능궈 ☆上관 없능궈능 亞능뒈훀 글애듴
ⓡ려듀九 싀풔숴훀 어릨 因仁川女中싁⑨들앜!! ㉯㉯납♡ ⌒⌒*`,
	},
}

func TestFrom(t *testing.T) {
	for _, test := range tests {
		t.Logf("Testing From for %s...", test.utf8)
		expect := []byte(test.utf8)
		got, err := From([]byte(test.cp949))
		if err != nil {
			t.Error(err)
		}
		if 0 != bytes.Compare(got, expect) {
			t.Errorf("\nexpect\t%v but,\ngot\t%v", expect, got)
		}
	}
}

func TestTo(t *testing.T) {
	for _, test := range tests {
		t.Logf("Testing To for %s...", test.utf8)
		expect := []byte(test.cp949)
		got, err := To([]byte(test.utf8))
		if err != nil {
			t.Error(err)
		}
		if 0 != bytes.Compare(got, expect) {
			t.Errorf("\nexpect\t%v but,\ngot\t%v", expect, got)
		}
	}
}

func (t lookupTable) Len() int {
	return len(t)
}

func (t lookupTable) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type lookupTableSortByCp949 struct{ lookupTable }

func (t lookupTableSortByCp949) Less(i, j int) bool {
	return t.lookupTable[i].cp949 < t.lookupTable[j].cp949
}

type lookupTableSortByUcs2 struct{ lookupTable }

func (t lookupTableSortByUcs2) Less(i, j int) bool {
	return t.lookupTable[i].ucs2 < t.lookupTable[j].ucs2
}

func TestLookupTableSorted(t *testing.T) {
	if !sort.IsSorted(lookupTableSortByCp949{fromTable}) {
		t.Error("fromTable is not sorted!")
	}
	t.Log("fromTable is sorted")

	if !sort.IsSorted(lookupTableSortByUcs2{toTable}) {
		t.Error("toTable is not sorted!")
	}
	t.Log("toTable is sorted")
}
