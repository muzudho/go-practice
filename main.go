package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muzudho/go-practice/exercise"
)

func main() {
	fmt.Print("📁exerciseフォルダ下の📄ファイル名が練習名です。練習名を入力してください　｜　例 strings　｜　例 exit　：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		practiceName := scanner.Text() // 1行ずつテキストを取得します

		if practiceName == "exit" {
			break // "exit"と入力されたらループを抜けます
		}

		fmt.Printf("練習名：%s\n", practiceName)

		switch practiceName {
		case "strings":
			exercise.Strings()
		case "fmt":
			exercise.Fmt()
		}

		fmt.Print("\n練習名を入力してください　｜　例 strings　｜　例 exit　：")
	}

}
