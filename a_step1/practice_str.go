package a_step1

import (
	"fmt"
	"strings"
	"unicode"
)

// SubRoutine - 練習１
func SubRoutine() {
	var s = "Hello, world!"
	fmt.Println(s)

	// 先頭の数文字を読むには（＾ｑ＾）？
	fmt.Println(s[0])   // なんか１文字取ったら数が出てくるぜ（＾ｑ＾）
	fmt.Println(s[0:1]) // こう書くと文字が出てくるぜ（＾ｑ＾）
	fmt.Println(s[0:2])
	fmt.Println(s[:2])
	fmt.Println(s[2:])

	// フォーマットも使ってみようぜ（＾ｑ＾）？
	fmt.Printf("%d\n", s[0]) // 数
	fmt.Printf("%c\n", s[0]) // %s ではなく %c

	var s2 = "日本語だったらどうなる（＾ｑ＾）？"
	fmt.Println(s2)

	// 先頭の数文字を読むには（＾ｑ＾）？
	fmt.Println(s2[0])   // なんか１文字取ったら数が出てくるぜ（＾ｑ＾）
	fmt.Println(s2[0:1]) // 文字化け
	fmt.Println(s2[0:2]) // 2bytes 文字化け
	fmt.Println(s2[0:3]) // 日
	fmt.Println(s2[0:4]) // 日と1byte文字化け。文字ごとのバイトサイズは分からんなあ
	fmt.Println(s2[:2])
	fmt.Println(s2[2:])

	// フォーマットも使ってみようぜ（＾ｑ＾）？
	fmt.Printf("%d\n", s2[0])
	fmt.Printf("%c\n", s2[0]) // %s ではなく %c。でもダメ文字化け

	// runeの配列にしろとのこと
	r3 := []rune(s2)

	fmt.Println("日本語を読みたいぜ（＾ｑ＾）！")
	// 先頭の数文字を読むには（＾ｑ＾）？
	fmt.Println(r3[0])         // 26085
	fmt.Println(string(r3[0])) // 日
	fmt.Println(string(r3[0:1]))
	fmt.Println(string(r3[0:2]))
	fmt.Println(string(r3[:2]))
	fmt.Println(string(r3[2:]))

	// フォーマットも使ってみようぜ（＾ｑ＾）？
	fmt.Printf("%d\n", r3[0])           // 26085
	fmt.Printf("%c\n", r3[0])           // 日
	fmt.Printf("%c\n", r3[0:3])         // [日 本 語]
	fmt.Printf("%s\n", string(r3[0:3])) // 日本語

	fmt.Println("runeの配列を回したらいいのかだぜ（＾ｑ＾）！？")
	// rangeで取り出すとrune単位で取り出せる。
	for _, c3 := range r3 {
		fmt.Println(string(c3))
	}
}

// SubRoutine2 - 練習２
func SubRoutine2() {
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

// SubRoutine3 - 練習３
func SubRoutine3() {
	var s1 = "apple banana cherry"
	fmt.Println(s1)

	res1 := strings.Index(s1, "banana")
	fmt.Println("Index: ", res1)
}
