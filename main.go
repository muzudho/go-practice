package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/muzudho/go-practice/exercise"
)

func main() {
	// // コマンドライン引数登録関数
	// func commandLineToPArgsMap(commandLine string) map[string]*string {
	// 	// フラグセットを作成（エラー時はプログラムを終了）
	// 	fs2 := flag.NewFlagSet("custom-args", flag.ExitOnError)	// 2. 新規フラグセットを作成（エラー時はプログラムを終了）
	// 	// コマンドライン引数名と、その値が入る変数へのポインターを紐づけるマップ
	// 	pArgsMap := make(map[string]*string)

	// 	// コマンドライン引数を登録し、後でその値が入る変数へのポインターを取得
	// 	pArgsMap["p"] = fs2.String("p", "", "Practice name. It is the file name under the 📁exercise folder.")

	// 	parameters := strings.Split(commandLine, " ") // コマンドライン引数をすべて取得
	// 	fs2.Parse(parameters[1:])     // コマンドライン引数の解析

	// 	return pArgsMap
	// }

	commandLine1 := strings.Join(os.Args, " ") // 1. コマンドラインを文字列として取得

	if commandLine1 == "exit" {
		return // "exit"と入力されたらプログラムを抜けます
	}

	onCommandLineEntered(commandLine1)

	fs1 := flag.NewFlagSet("main-args", flag.ExitOnError) // 1. 新規フラグセットを作成（エラー時はプログラムを終了）

	//fs1 := flag.CommandLine                                                                              // 2. コマンドラインに紐づいたフラグセットを取得
	pArgsMap := make(map[string]*string)                                                                 // 3. コマンドライン引数名と、その値が入る変数へのポインターを紐づけるマップ
	pArgsMap["p"] = fs1.String("p", "", "Program name. It is the file name under the 📁exercise folder.") // 4. コマンドライン引数を登録し、後でその値が入る変数へのポインターを取得

	parameters := os.Args[1:]            // 5. コマンドライン引数をすべて取得
	fs1.Parse(parameters)                // 6. コマンドライン引数の解析
	fmt.Printf("p=%s\n", *pArgsMap["p"]) // ヌルを指していれば、空文字列になるだけ。問題ない。
	// 7. （あれば）必須のコマンドライン引数の確認

	fmt.Print("Please enter the program name ｜ e.g. hello ｜ e.g. exit ：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		commandLine2 := scanner.Text() // 1行ずつテキストを取得します

		if commandLine2 == "exit" {
			break // "exit"と入力されたらループを抜けます
		}

		onCommandLineEntered(commandLine2)

		// practiceName を空白でスプリットし、最初の要素を取得します
		tokens := strings.Split(commandLine2, " ")
		practiceName := tokens[0]

		fmt.Printf("練習名：%s\n", practiceName)

		switch practiceName {
		case "echo_stdio":
			exercise.EchoStdio()
		case "echo_proxy":
			// ```
			// echo_proxy Z:/muzudho-github.com/muzudho/go-practice/go-practice.exe
			// ```
			exercise.EchoProxy(tokens[1])
		case "fmt":
			exercise.Fmt()
		case "hello":
			exercise.Hello()
		case "strings":
			exercise.Strings()
		}

		fmt.Print("\n練習名を入力してください　｜　例 strings　｜　例 exit　：")
	}
}

func onCommandLineEntered(commandLine string) {
	fmt.Printf("Command line entered: [%s]\n", commandLine)
}
