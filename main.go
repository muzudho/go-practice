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
	// // コマンドラインを文字列として取得
	// fullCmdLine := strings.Join(os.Args, " ")
	// fmt.Printf("Full command line: [%s]\n", fullCmdLine)

	// // コマンドライン引数登録関数
	// func commandLineToPArgsMap(commandLine string) map[string]*string {
	// 	// フラグセットを作成（エラー時はプログラムを終了）
	// 	fs2 := flag.NewFlagSet("custom-args", flag.ExitOnError)
	// 	// コマンドライン引数名と、その値が入る変数へのポインターを紐づけるマップ
	// 	pArgsMap := make(map[string]*string)

	// 	// コマンドライン引数を登録し、後でその値が入る変数へのポインターを取得
	// 	pArgsMap["p"] = fs2.String("p", "", "Practice name. It is the file name under the 📁exercise folder.")

	// 	parameters := strings.Split(commandLine, " ") // コマンドライン引数をすべて取得
	// 	fs2.Parse(parameters[1:])     // コマンドライン引数の解析

	// 	return pArgsMap
	// }
	fs1 := flag.CommandLine              // ← これでコマンドラインに紐づいたフラグセットをゲット！
	pArgsMap := make(map[string]*string) // コマンドライン引数名と、その値が入る変数へのポインターを紐づけるマップ

	// コマンドライン引数を登録し、後でその値が入る変数へのポインターを取得
	pArgsMap["p"] = fs1.String("p", "", "Practice name. It is the file name under the 📁exercise folder.")

	parameters := os.Args[1:] // コマンドライン引数をすべて取得
	fs1.Parse(parameters)     // コマンドライン引数の解析

	// デバッグ出力
	fmt.Printf("p=%s\n", *pArgsMap["p"]) // ヌルを指していれば、空文字列になるだけ。問題ない。

	// 📁exerciseフォルダ下の📄ファイル名が練習名です。中には引数が必要なものもあります。練習名を入力してください　｜　例 strings　｜　例 exit　：
	// を英語で：
	fmt.Print("Practice name is the file name under the 📁exercise folder. Please enter the practice name ｜ e.g. strings ｜ e.g. exit ：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		commandLine := scanner.Text() // 1行ずつテキストを取得します

		if commandLine == "exit" {
			break // "exit"と入力されたらループを抜けます
		}

		// practiceName を空白でスプリットし、最初の要素を取得します
		tokens := strings.Split(commandLine, " ")
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
		case "strings":
			exercise.Strings()
		}

		fmt.Print("\n練習名を入力してください　｜　例 strings　｜　例 exit　：")
	}

}
