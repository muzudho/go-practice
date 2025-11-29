package exercise

import (
	"bufio"
	"os"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var next = "BCDEFGHIJKLMNOPQRSTUVWXYZAbcdefghijklmnopqrstuvwxyza"

// Bufio - bufioの練習
func Bufio() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		command := scanner.Text() // 1行ずつテキストを取得します

		if 0 < len(command) {
			one := command[0:1] // 1文字目を取得します

			index := strings.Index(alphabet, one) // その文字が［alphabet］文字列の何番目にあるか調べます
			if index != -1 {
				next := next[index : index+1] // ［next］文字列の同じ位置の文字を取得します
				os.Stdout.WriteString(next)
			}
		}
	}
}
