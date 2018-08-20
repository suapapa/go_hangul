
# About hangul

hangul is a set of handy tools for manipulate korean character in Go language.

[![GoDoc](https://godoc.org/github.com/suapapa/go_hangul?status.svg)](https://godoc.org/github.com/suapapa/go_hangul)
[![Build Status](https://travis-ci.org/suapapa/go_hangul.png?branch=master)](https://travis-ci.org/suapapa/go_hangul)

## Example

    package main

    import (
            "fmt"
            hangul "github.com/suapapa/go_hangul"
    )

    func main() {
            var i, m, f, ci, cm, cf rune
            var c int
            for _, r := range "맥도날드" {
                    // Storke count
                    c = hangul.Stroke(r)
                    // Split to three elements
                    i, m, f = hangul.Split(r)
                    // Convert between jamo and compatibility-jamo
                    ci = hangul.CompatJamo(i)
                    cm = hangul.CompatJamo(m)
                    cf = hangul.CompatJamo(f)

                    fmt.Printf("%c %d %c(%v) %c(%v) %c(%v)\n", r, c, ci, i, cm, m, cf, f)
            }
            fmt.Println("")

            fmt.Println(
                hangul.EndsWithConsonant("강")) // true
            fmt.Println(
                hangul.EndsWithConsonant("그")) // false
            fmt.Println(
                hangul.AppendPostposition("강", "이", "가")) // "강이"
            fmt.Println(
                hangul.AppendPostposition("물고기", "은", "는")) // "물고기는"
    }

# Installation

    $ go get github.com/suapapa/go_hangul

# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
