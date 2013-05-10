// Copyright 2013, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hanja

import (
	"testing"
)

func TestTable(t *testing.T) {
	hanja := "大韓民國은 民主共和國이다."
	hangul := "대한민국은 민주공화국이다."

	if Convert(hanja) != hangul {
		t.Errorf("expect %s but, got %s", hangul, Convert(hanja))
	}
}
