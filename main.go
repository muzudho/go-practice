package main

import "fmt"

func main() {
	var s = "Hello, world!"
	fmt.Println(s)

	// 先頭の数文字を読むには（＾ｑ＾）？
	fmt.Println(s[0])   // なんか１文字取ったら数が出てくるぜ（＾ｑ＾）
	fmt.Println(s[0:1]) // こう書くと文字が出てくるぜ（＾ｑ＾）
	fmt.Println(s[0:2])
	fmt.Println(s[:2])
	fmt.Println(s[2:])

	// フォーマットも使ってみようぜ（＾ｑ＾）？
	fmt.Printf("%d\n", s[0])
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
}
