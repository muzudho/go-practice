package exercise

import (
	"fmt"
	"strings"
)

// Strings - 文字列の練習
func Strings() {
	var s1 = "apple banana cherry"
	fmt.Println(s1) // 文字列表示

	res1 := strings.Index(s1, "banana")
	fmt.Println("banana's index: ", res1) // bananaの位置

	res2 := strings.HasPrefix(s1, "apple") // 先頭がappleか？
	fmt.Println("apple is first: ", res2)
}
