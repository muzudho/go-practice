package exercise

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// String - 文字列の練習
func String(targetString string) {

	var s = targetString // 変数名を短縮

	fmt.Printf("The number of letters in \"%s\" is %d.\n", s, utf8.RuneCountInString(s)) // 文字数を出力。RuneはUnicode文字。

	var s1 = "apple banana cherry"
	fmt.Println(s1) // 文字列表示

	res1 := strings.Index(s1, "banana")
	fmt.Println("banana's index: ", res1) // bananaの位置

	res2 := strings.HasPrefix(s1, "apple") // 先頭がappleか？
	fmt.Println("apple is first: ", res2)
}
