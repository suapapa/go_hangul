
# About hangul

hangul is a set of handy tools for manipulate korean character which can:

- Convert between jamo and compatibility-jamo.
- Split a character to it's three elements.
- Stroke count of a jamo.

## TODO

- Romanize.
- Lexycal Analize.


# Documentation

## Prerequisites

[Install Go][].

## Installation

    $ go get github.com/suapapa/hangul

## General Documentation

Use `go doc` to vew the documentation for hangul

    go doc github.com/suapapa/hangul

Or alternatively, use a godoc http server

    godoc -http=:6060

and visit [the Godoc URL][]

# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.

# References

- [한국어 위키낱말사전][1]
    * [한국어 개정 로마자 표기법][1-1]

[install go]: http://golang.org/install.html "Install Go"
[the godoc url]: http://localhost:6060/pkg/github.com/suapapa/hangul/ "the Godoc URL"

[1]: http://ko.wiktionary.org
[1-1]: http://ko.wiktionary.org/wiki/%EC%9C%84%ED%82%A4%EB%82%B1%EB%A7%90%EC%82%AC%EC%A0%84:%EB%A1%9C%EB%A7%88%EC%9E%90_%ED%91%9C%EA%B8%B0%EB%B2%95/%ED%95%9C%EA%B8%80
