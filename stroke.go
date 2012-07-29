// Copyright 2012, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hangul

var strokes = map[rune]int{
	G:   1,
	N:   1,
	D:   2,
	L:   3,
	M:   3,
	B:   4,
	S:   2,
	ZS:  1,
	J:   3,
	C:   4,
	K:   2,
	T:   3,
	P:   4,
	H:   3,
	A:   2,
	YA:  3,
	EO:  2,
	YEO: 3,
	O:   2,
	YO:  3,
	U:   2,
	YU:  3,
	EU:  1,
	I:   1,
}

// Get stroke count of jamo.
func Stroke(r rune) (c int) {
	if 0xAC00 <= r && r <= 0xD7A3 {
		i, m, f := Split(r)
		c += Stroke(i)
		c += Stroke(m)
		c += Stroke(f)
		return
	}

	jm := CompatJamo(r)
	if es, ok := SplitMultiElement(jm); ok {
		for _, e := range es {
			c += strokes[e]
		}
	} else {
		c = strokes[jm]
	}
	return
}
