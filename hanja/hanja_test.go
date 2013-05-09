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
