// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hangul

/*  Filename:    hangul_test.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-07-16 17:16:58.048792 +0900 KST
 *  Description: Main test file for hangul
 */

import (
	"testing"
)

func TestIdx(t *testing.T) {
	if leadIdx(Lead(T)) != 16 {
		t.Errorf("Got %v, expect %v", leadIdx(Lead(T)), 16) // 양
	}
}

func TestStroke(t *testing.T) {
	var i int
	var ss = []int{5, 5, 3, 0, 4, 5}
	for _, c := range "세상아 안녕" {
		if ss[i] != Stroke(c) {
			t.Errorf("%c Expected %d got %d", c, ss[i], Stroke(c))
		}
		i += 1
	}

	if c := Stroke(JJ); c != 6 {
		t.Errorf("Unexpected count %d for JJ", c)
	}
	if c := Stroke(YAE); c != 4 {
		t.Errorf("Unexpected count %d for YAE", c)
	}
	if c := Stroke(WAE); c != 5 {
		t.Errorf("Unexpected count %d for WAE", c)
	}
}

func TestJoin(t *testing.T) {
	if c := Join(LEAD_S, MEDIAL_EO, 0); c != 0xC11C {
		t.Errorf("Got %c(%v), expect %c(%v)", c, c, 0xC11C, 0xC11C) // 서
	}
	if c := Join(LEAD_ZS, MEDIAL_U, TAIL_L); c != 0xC6B8 {
		t.Errorf("Got %c, expect %c", c, 0xC6B8) // 울
	}
	if c := Join(LEAD_P, MEDIAL_YEO, TAIL_NG); c != 0xD3C9 {
		t.Errorf("Got %c, expect %c", c, 0xD3C9) // 평
	}
	if c := Join(LEAD_ZS, MEDIAL_YA, TAIL_NG); c != 0xC591 {
		t.Errorf("Got %c, expect %c", c, 0xC11C) // 양
	}
}

func TestSplit(t *testing.T) {
	type hangul struct {
		i, m, f rune
	}

	var splitted = []hangul{
		{J, A, 0},    // 자
		{M, O, 0},    // 모
		{H, A, N},    // 한
		{G, EU, L},   // 글
		{ZS, A, N},   // 안
		{N, YEO, NG}, // 녕
	}

	var idx int
	for _, c := range "자모한글안녕" {
		exp := splitted[idx]
		i, m, f := SplitCompat(c)
		switch {
		case i != exp.i:
			t.Errorf("%c-l: expected %c, got %", c, exp.i, i)
			return
		case m != exp.m:
			t.Errorf("%c-m: expected %c, got %c", c, exp.m, m)
			return
		case f != exp.f:
			t.Errorf("%c-t: expected %c, got %c", c, exp.f, f)
			return

		}
		idx += 1
	}
}

func TestMultiElements(t *testing.T) {
	if es, ok := SplitMultiElement(GG); ok {
		if len(es) != 2 {
			t.Errorf("GG != G, G??\n")
		}
		if es[0] != G || es[1] != G {
			t.Errorf("%v\n", es)
		}
	} else {
		t.Errorf("Failed to get multielements\n")
	}

	if _, ok := SplitMultiElement(G); ok {
		t.Errorf("G is not multi element\n")
	}
}

func TestComaptJamo(t *testing.T) {
	if c := CompatJamo(LEAD_G); c != G {
		t.Errorf("Failed to convert to comaptibility jamo! ")
		t.Errorf("expected %v but, got %v\n", G, c)
	}
	if c := CompatJamo(TAIL_H); c != H {
		t.Errorf("Failed to convert to comaptibility jamo! ")
		t.Errorf("expected %v but, got %v\n", H, c)
	}
	if c := CompatJamo(MEDIAL_A); c != A {
		t.Errorf("Failed to convert to comaptibility jamo! ")
		t.Errorf("expected %v but, got %v\n", A, c)
	}
	if c := CompatJamo(MEDIAL_I); c != I {
		t.Errorf("Failed to convert to comaptibility jamo! ")
		t.Errorf("expected %v but, got %v\n", I, c)
	}
}

func TestJamoConstants(t *testing.T) {
	if H != 0x314E {
		t.Errorf("Last Jaeum sholud be 0x314E."+
			" not 0x%04x\n", H)
	}

	if I != 0x3163 {
		t.Errorf("Last Moeum sholud be 0x3163."+
			" not 0x%04x\n", I)
	}

	if LEAD_H != 0x1112 {
		t.Errorf("Last Lead sholud be 0x1112."+
			" not 0x%04x\n", LEAD_H)
	}

	if MEDIAL_I != 0x1175 {
		t.Errorf("Last Medial sholud be 0x1175."+
			" not 0x%04x\n", MEDIAL_I)
	}

	if TAIL_H != 0x11C2 {
		t.Errorf("Last Tail sholud be 0x11C3."+
			" not 0x%04x\n", TAIL_H)
	}
}
