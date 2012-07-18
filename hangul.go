// Copyright 2012, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hangul

import (
//	"log"
//	"unicode"
//	"unicode/utf16"
)

// Check Given rune is Hangul
func IsHangul(r rune) bool {
	switch {
	case 44032 <= r && r <= 0xD7A3:
		return true
	case IsJaeum(r):
		return true
	case IsMoeum(r):
		return true
	}

	return false
}

// Convert NFD to NFC
func Join(l, m, t rune) rune {
	// Convert if given rune is compatibility jamo
	l = Lead(l)
	m = Medial(m)
	t = Tail(t)
	c := leadIdx(l)*588 + medialIdx(m)*28 + tailIdx(t) + 44032
	return rune(c)
}

// Convert NFC to NFD
func Split(c rune) (l, m, t rune) {
	t = (c - 44032) % 28
	m = ((c - 44032 - t) % 588) / 28
	l = (c - 44032) / 588

	l += LEAD_BASE
	m += MEDIAL_BASE
	if t != 0 {
		t += TAIL_BASE
	}
	return
}

// Split and got l, m, t in compatibility jamo
func SplitCompat(c rune) (l, m, t rune) {
	l, m, t = Split(c)
	l = CompatJamo(l)
	m = CompatJamo(m)
	t = CompatJamo(t)
	return
}

func ConJoin(s *string) {
}

func DisJoint(s *string) {
}
