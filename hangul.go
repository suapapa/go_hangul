// Copyright 2012, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hangul provide handy tools for manipulate korean character:
//    - Provide shorthands for korean consonants and vowels
//    - Convert between jamo and compatibility-jamo
//    - Split a character to it's three elements
//    - Split multi element
//    - Stroke count
package hangul

// IsHangul checks given rune is Hangul
func IsHangul(r rune) bool {
	switch {
	case 0xAC00 <= r && r <= 0xD7A3:
		return true
	case IsJaeum(r):
		return true
	case IsMoeum(r):
		return true
	}

	return false
}

// Join converts NFD to NFC
func Join(l, m, t rune) rune {
	// Convert if given rune is compatibility jamo
	li, ok := leadIdx(Lead(l))
	if !ok {
		return rune(0xFFFD)
	}

	mi, ok := medialIdx(Medial(m))
	if !ok {
		return rune(0xFFFD)
	}

	ti, ok := tailIdx(Tail(t))
	if !ok {
		return rune(0xFFFD)
	}

	return rune(0xAC00 + (li*21+mi)*28 + ti)
}

// Split converts NFC to NFD
func Split(c rune) (l, m, t rune) {
	t = (c - 0xAC00) % 28
	m = ((c - 0xAC00 - t) % 588) / 28
	l = (c - 0xAC00) / 588

	l += LEAD_BASE
	m += MEDIAL_BASE
	if t != 0 {
		t += TAIL_BASE - 1
	}

	return
}

// SplitCompat splits and returns l, m, t in compatibility jamo
func SplitCompat(c rune) (l, m, t rune) {
	l, m, t = Split(c)
	l = CompatJamo(l)
	m = CompatJamo(m)
	t = CompatJamo(t)
	return
}

// EndsWithConsonant returns true if the given word ends with
// consonant(종성), false otherwise.
func EndsWithConsonant(word string) bool {
	if LastConsonant(word) == 0 {
		return false
	}

	return true
}

// LastConsonant returns last consonant(종성).
// It returns 0 if last consonant not exists.
func LastConsonant(word string) rune {
	if word == "" {
		return 0
	}
	runes := []rune(word)
	_, _, t := Split(runes[len(runes)-1])
	return t
}

// AppendPostposition returns word with postposition(조사).
// The 'with' will be appended to a word that ends with consonant, and
// the 'without' will be appended to a word that does not end with
// consonant.
func AppendPostposition(word, with, without string) string {
	lastTail := LastConsonant(word)
	if with == "으로" && lastTail == TAIL_L {
		return word + without // 로
	}
	if lastTail != 0 {
		return word + with
	}
	return word + without
}
