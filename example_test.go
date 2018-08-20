package hangul_test

import (
	"fmt"

	hangul "github.com/suapapa/go_hangul"
)

func ExampleAppendPostposition() {
	fmt.Println(hangul.AppendPostposition("강", "이", "가"))
	fmt.Println(hangul.AppendPostposition("물고기", "은", "는"))
	fmt.Println(hangul.AppendPostposition("영철", "이랑", "랑"))
	fmt.Println(hangul.AppendPostposition("순희", "이랑", "랑"))
	fmt.Println(hangul.AppendPostposition("마을", "으로", "로"))
	// Output:
	// 강이
	// 물고기는
	// 영철이랑
	// 순희랑
	// 마을로
}

func ExampleLastConsonant() {
	fmt.Println(hangul.LastConsonant("강"))
	fmt.Println(hangul.LastConsonant("물고기"))
	// Output:
	// 4540
	// 0
}
