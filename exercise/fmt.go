package exercise

import (
	"fmt"
	"strings"
	"unicode"
)

// Fmt - 文字列出力の練習
func Fmt() {
	// runeの配列にしろとのこと
	r1 := []rune("ハローワールドだぜ（＾ｑ＾）")

	fmt.Printf("%s", string(r1))

	// 配列アクセスしたいぜ（＾ｑ＾）
	fmt.Printf("%s\n", string(r1[0]))
	fmt.Printf("%s\n", string(r1[1]))
	fmt.Printf("%s\n", string(r1[2]))
	fmt.Printf("%s\n", string(r1[3:5]))
	fmt.Printf("%s\n", string(r1[3:]))
	fmt.Printf("%s\n", string(r1[:5]))

	// 先頭から次のスペースまで読むって できんの（＾ｑ＾）？
	r2 := []rune("apple バナナ Cherry")
	res2 := strings.IndexRune(string(r2), ' ')
	fmt.Printf("res2=%d\n", res2)
	// 先頭から次のスペースまでの文字を読みたいぜ（＾ｑ＾）
	fmt.Printf("0:2=%s\n", string(r2[0:res2])) // apple

	// 読み取った文字が、アルファベットか、数か 区別する方法あんの（＾ｑ＾）？
	r3 := []rune("体重は90kgぐらいかだぜ（＾ｑ＾）？")
	res3 := unicode.IsLetter(r3[3])
	fmt.Printf("%c is letter? = %t\n", r3[3], res3) // boolean は %t
	res3 = unicode.IsNumber(r3[3])
	fmt.Printf("%c is number? = %t\n", r3[3], res3)
}
