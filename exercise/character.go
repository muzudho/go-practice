package exercise

import (
	"fmt"
	"unicode/utf8"
)

// Character - 文字操作の練習
func Character(targetString string) {

	var s = targetString // 変数名を短縮

	fmt.Printf("The number of letters in \"%s\" is %d.\n", s, utf8.RuneCountInString(s)) // 文字数を出力。RuneはUnicode文字。
}
