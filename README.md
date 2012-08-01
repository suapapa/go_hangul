
[install go]: http://golang.org/install.html "Install Go"

# About hangul

![logo](https://lh5.googleusercontent.com/-yblxhHfOiXw/UAzP9_3B0FI/AAAAAAAAA74/0nKCplLb9Ck/s615/IMG_20120723_131321-1.jpg)

hangul is a set of handy tools for manipulate korean character.

## Example

    package main

    import (
            "fmt"
            "github.com/suapapa/go_hangul"
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
    }

# Documentation

## Prerequisites

[Install Go][]

## Installation

    $ go get github.com/suapapa/go_hangul

## General Documentation

Use `go doc` to vew the documentation for hangul

    go doc github.com/suapapa/go_hangul

Or alternatively, refer [go.pkgdoc.org](http://go.pkgdoc.org/github.com/suapapa/go_hangul)


# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
