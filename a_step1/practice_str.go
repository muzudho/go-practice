package a_step1

import "fmt"

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

// SubRoutine - 練習２
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
}
