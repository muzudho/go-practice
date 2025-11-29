package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muzudho/go-practice/a_step1"
)

func main() {
	fmt.Print("練習名を入力してください　｜　例 hello　｜　例 exit　：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		practiceName := scanner.Text() // 1行ずつテキストを取得します

		if practiceName == "exit" {
			break // "exit"と入力されたらループを抜けます
		}

		fmt.Printf("練習名：%s", practiceName)

		if practiceName == "hello" {
			// a_step1.SubRoutine()
			// a_step1.SubRoutine2()
			a_step1.SubRoutine3()
		}

		fmt.Print("\n練習名を入力してください　｜　例 hello　｜　例 exit　：")
	}

}
