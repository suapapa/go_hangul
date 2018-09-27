// Copyright 2012, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hangul

// IsLead checks given rune is lead consonant
func IsLead(r rune) bool {
	if LeadG <= r && r <= LeadH {
		return true
	}
	return false
}

// IsMedial checks given rune is medial vowel
func IsMedial(r rune) bool {
	if MedialA <= r && r <= MedialI {
		return true
	}
	return false
}

// IsTail checks given rune is tail consonant
func IsTail(r rune) bool {
	if TailG <= r && r <= TailH {
		return true
	}
	return false
}

// IsJaeum checks given rune is Hangul Jaeum
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

// IsMoeum checks given rune is Hangul Moeum
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

// SplitMultiElement splits multi-element compatibility jamo
func SplitMultiElement(r rune) ([]rune, bool) {
	r = CompatJamo(r)
	es, ok := multiElements[r]
	return es, ok
}

var toCompatJamo = map[rune]rune{
	LeadG:  G,
	TailG:  G,
	LeadGG: GG,
	TailGG: GG,
	TailGS: GS,
	LeadN:  N,
	TailN:  N,
	TailNJ: NJ,
	TailNH: NH,
	LeadD:  D,
	TailD:  D,
	LeadDD: DD,
	LeadR:  L,
	TailL:  L,
	TailLG: LG,
	TailLM: LM,
	TailLB: LB,
	TailLS: LS,
	TailLT: LT,
	TailLP: LP,
	TailLH: LH,
	LeadM:  M,
	TailM:  M,
	LeadB:  B,
	TailB:  B,
	LeadBB: BB,
	TailBS: BS,
	LeadS:  S,
	TailS:  S,
	LeadSS: SS,
	TailSS: SS,
	LeadZS: ZS,
	TailNG: ZS,
	LeadJ:  J,
	TailJ:  J,
	LeadJJ: JJ,
	LeadC:  C,
	TailC:  C,
	LeadK:  K,
	TailK:  K,
	LeadT:  T,
	TailT:  T,
	LeadP:  P,
	TailP:  P,
	LeadH:  H,
	TailH:  H,
}

// CompatJamo converts lead, medial, tail to compatibility jamo
func CompatJamo(r rune) rune {
	switch {
	case G <= r && r <= H:
		return r
	case A <= r && r <= I:
		return r
	case MedialA <= r && r <= MedialI:
		return r - medialBase + A
	}
	if c, ok := toCompatJamo[r]; ok {
		return c
	}

	return 0
}

var toLead = map[rune]rune{
	G:  LeadG,
	GG: LeadGG,
	N:  LeadN,
	D:  LeadD,
	DD: LeadDD,
	L:  LeadR,
	M:  LeadM,
	B:  LeadB,
	BB: LeadBB,
	S:  LeadS,
	SS: LeadSS,
	ZS: LeadZS,
	J:  LeadJ,
	JJ: LeadJJ,
	C:  LeadC,
	K:  LeadK,
	T:  LeadT,
	P:  LeadP,
	H:  LeadH,
}

// Lead converts compatibility jaeum to corresponding lead consonant
func Lead(c rune) rune {
	if LeadG <= c && c <= LeadH {
		return c
	}
	if l, ok := toLead[c]; ok {
		return l
	}

	return 0
}

// Medial converts compatibility moeum to corresponding medial vowel
func Medial(c rune) rune {
	switch {
	case MedialA <= c && c <= MedialI:
		return c
	case A <= c && c <= I:
		return c - A + medialBase
	}

	return 0
}

var toTail = map[rune]rune{
	G:  TailG,
	GG: TailGG,
	GS: TailGS,
	N:  TailN,
	NJ: TailNJ,
	NH: TailNH,
	D:  TailD,
	L:  TailL,
	LG: TailLG,
	LM: TailLM,
	LB: TailLB,
	LS: TailLS,
	LT: TailLT,
	LP: TailLP,
	LH: TailLH,
	M:  TailM,
	B:  TailB,
	BS: TailBS,
	S:  TailS,
	SS: TailSS,
	ZS: TailNG,
	J:  TailJ,
	C:  TailC,
	K:  TailK,
	T:  TailT,
	P:  TailP,
	H:  TailH,
}

// Tail converts compatibility jaeum to corresponding tail consonant
func Tail(c rune) rune {
	if TailG <= c && c <= TailH {
		return c
	}
	if t, ok := toTail[c]; ok {
		return t
	}

	return 0
}

func leadIdx(l rune) (int, bool) {
	i := int(l) - leadBase
	if 0 > i || i > maxLeadIdx {
		return 0, false
	}
	return i, true
}

func medialIdx(v rune) (int, bool) {
	i := int(v) - medialBase
	if 0 > i || i > maxMedialIdx {
		return 0, false
	}
	return i, true
}

func tailIdx(t rune) (int, bool) {
	if t == 0 {
		// A hangul syllable can have no tail consonent.
		return 0, true
	}
	i := int(t) - tailBase
	if 0 > i || i > maxTailIdx {
		return 0, false
	}
	return i + 1, true
}
