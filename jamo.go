// Copyright 2012, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hangul

import (
	"log"
)

// Check Given rune is Lead consonant
func IsLead(r rune) bool {
	if LEAD_G <= r && r <= LEAD_H {
		return true
	}
	return false
}

// Check Given rune is Medial vowel
func IsMedial(r rune) bool {
	if MEDIAL_A <= r && r <= MEDIAL_I {
		return true
	}
	return false
}

// Check Given rune is Tail consonant
func IsTail(r rune) bool {
	if TAIL_G <= r && r <= TAIL_H {
		return true
	}
	return false
}

// Check Given rune is Hangul Jaeum
func IsJaeum(r rune) bool {
	switch {
	case G <= r && r <= H:
		return true
	case IsLead(r):
		return true
	case IsTail(r):
		return true
	}
	return false
}

// Check Given rune is Hangul Moeum
func IsMoeum(r rune) bool {
	switch {
	case A <= r && r <= I:
		return true
	case IsMedial(r):
		return true
	}
	return false
}

var multiElements = map[rune][]rune{
	GG:  []rune{G, G},
	GS:  []rune{G, S},
	NJ:  []rune{N, J},
	NH:  []rune{N, H},
	DD:  []rune{D, D},
	LG:  []rune{L, G},
	LM:  []rune{L, M},
	LB:  []rune{L, B},
	LS:  []rune{L, S},
	LT:  []rune{L, T},
	LP:  []rune{L, P},
	LH:  []rune{L, H},
	BB:  []rune{B, B},
	BS:  []rune{B, S},
	SS:  []rune{S, S},
	JJ:  []rune{J, J},
	AE:  []rune{A, I},
	E:   []rune{EO, I},
	YAE: []rune{YA, I},
	YE:  []rune{YEO, I},
	WA:  []rune{O, A},
	WAE: []rune{O, A, I},
	OE:  []rune{O, I},
	WEO: []rune{U, EO},
	WE:  []rune{U, E},
	WI:  []rune{U, I},
	YI:  []rune{EU, I},
}

// Split multi-element compatibility jamo
func SplitMultiElement(r rune) ([]rune, bool) {
	r = CompatJamo(r)
	es, ok := multiElements[r]
	return es, ok
}

var toCompatJamo = map[rune]rune{
	LEAD_G:  G,
	TAIL_G:  G,
	LEAD_GG: GG,
	TAIL_GG: GG,
	TAIL_GS: GS,
	LEAD_N:  N,
	TAIL_N:  N,
	TAIL_NJ: NJ,
	TAIL_NH: NH,
	LEAD_D:  D,
	TAIL_D:  D,
	LEAD_DD: DD,
	LEAD_R:  L,
	TAIL_L:  L,
	TAIL_LG: LG,
	TAIL_LM: LM,
	TAIL_LB: LB,
	TAIL_LS: LS,
	TAIL_LT: LT,
	TAIL_LP: LP,
	TAIL_LH: LH,
	LEAD_M:  M,
	TAIL_M:  M,
	LEAD_B:  B,
	TAIL_B:  B,
	LEAD_BB: BB,
	TAIL_BS: BS,
	LEAD_S:  S,
	TAIL_S:  S,
	LEAD_SS: SS,
	TAIL_SS: SS,
	LEAD_ZS: ZS,
	TAIL_NG: ZS,
	LEAD_J:  J,
	TAIL_J:  J,
	LEAD_JJ: JJ,
	LEAD_C:  C,
	TAIL_C:  C,
	LEAD_K:  K,
	TAIL_K:  K,
	LEAD_T:  T,
	TAIL_T:  T,
	LEAD_P:  P,
	TAIL_P:  P,
	LEAD_H:  H,
	TAIL_H:  H,
}

// Convert lead, medial, tail to compatibility jamo
func CompatJamo(r rune) rune {
	switch {
	case G <= r && r <= H:
		return r
	case A <= r && r <= I:
		return r
	case MEDIAL_A <= r && r <= MEDIAL_I:
		return r - MEDIAL_BASE + A
	}
	if c, ok := toCompatJamo[r]; ok {
		return c
	}

	return 0
}

var toLead = map[rune]rune{
	G:  LEAD_G,
	GG: LEAD_GG,
	N:  LEAD_N,
	D:  LEAD_D,
	DD: LEAD_DD,
	L:  LEAD_R,
	M:  LEAD_M,
	B:  LEAD_B,
	BB: LEAD_BB,
	S:  LEAD_S,
	SS: LEAD_SS,
	ZS: LEAD_ZS,
	J:  LEAD_J,
	JJ: LEAD_JJ,
	C:  LEAD_C,
	K:  LEAD_K,
	T:  LEAD_T,
	P:  LEAD_P,
	H:  LEAD_H,
}

// Convert compatibility jaeum to corresponding lead consonant
func Lead(c rune) rune {
	if LEAD_G <= c && c <= LEAD_H {
		return c
	}
	if l, ok := toLead[c]; ok {
		return l
	}

	return 0
}

// Convert compatibility moeum to corresponding medial vowel
func Medial(c rune) rune {
	switch {
	case MEDIAL_A <= c && c <= MEDIAL_I:
		return c
	case A <= c && c <= I:
		return c - A + MEDIAL_BASE
	}

	return 0
}

var toTail = map[rune]rune{
	G:  TAIL_G,
	GG: TAIL_GG,
	GS: TAIL_GS,
	N:  TAIL_N,
	NJ: TAIL_NJ,
	NH: TAIL_NH,
	D:  TAIL_D,
	L:  TAIL_L,
	LG: TAIL_LG,
	LM: TAIL_LM,
	LB: TAIL_LB,
	LS: TAIL_LS,
	LT: TAIL_LT,
	LP: TAIL_LP,
	LH: TAIL_LH,
	M:  TAIL_M,
	B:  TAIL_B,
	BS: TAIL_BS,
	S:  TAIL_S,
	SS: TAIL_SS,
	ZS: TAIL_NG,
	J:  TAIL_J,
	C:  TAIL_C,
	K:  TAIL_K,
	T:  TAIL_T,
	P:  TAIL_P,
	H:  TAIL_H,
}

// Convert compatibility jaeum to corresponding tail consonant
func Tail(c rune) rune {
	if TAIL_G <= c && c <= TAIL_H {
		return c
	}
	if t, ok := toTail[c]; ok {
		return t
	}

	return 0
}

func leadIdx(l rune) int {
	i := int(l) - LEAD_BASE
	if 0 > i || i > MAX_LEAD_IDX {
		log.Fatalln("given %v isn't right lead character", l)
	}
	return i
}

func medialIdx(v rune) int {
	i := int(v) - MEDIAL_BASE
	if 0 > i || i > MAX_MEDIAL_IDX {
		log.Fatalln("given %v isn't right lead character", v)
	}
	return i
}

func tailIdx(t rune) int {
	if t == 0 {
		// A hangul syllable cat have no tail consonent.
		return 0
	}
	i := int(t) - TAIL_BASE
	if 0 > i || i > MAX_TAIL_IDX {
		log.Fatalln("given %v isn't right lead character", t)
	}
	return i + 1
}
