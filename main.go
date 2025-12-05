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
	commandLine1 := strings.Join(os.Args, " ") // 1. コマンドラインを文字列として取得

	if commandLine1 == "exit" {
		return // "exit"と入力されたらプログラムを抜けます
	}

	commandName, pArgsMap := parseCommandLine(commandLine1)
	fmt.Printf("commandName=%s, p=%s\n", commandName, *pArgsMap["p"]) // ちゃんとマッピングできたか確認。ヌルを指していれば、空文字列になるだけ。問題ない。

	fmt.Print("Please enter the program name ｜ e.g. hello ｜ e.g. exit ：")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 標準入力を読込みます
		commandLine2 := scanner.Text() // 1行ずつテキストを取得します

		if commandLine2 == "exit" {
			break // "exit"と入力されたらループを抜けます
		}

		commandName, pArgsMap = parseCommandLine(commandLine2)
		fmt.Printf("commandName=%s, p=%s\n", commandName, *pArgsMap["p"]) // ちゃんとマッピングできたか確認。ヌルを指していれば、空文字列になるだけ。問題ない。

		switch *pArgsMap["p"] {
		case "echo_stdio":
			exercise.EchoStdio()
		case "echo_proxy":
			// ```
			// echo_proxy -f Z:/muzudho-github.com/muzudho/go-practice/go-practice.exe
			// ```
			exercise.EchoProxy(*pArgsMap["f"])
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

func parseCommandLine(commandLine string) (string, map[string]*string) {
	fmt.Printf("Command line entered: [%s]\n", commandLine)

	// コマンドラインを半角空白で区切る
	tokens := strings.Split(commandLine, " ")

	fs1 := flag.NewFlagSet("main-args", flag.ExitOnError) // 1. 引数のマッピング（FlagSet）を作成（エラー時はプログラムを終了）

	pArgsMap := make(map[string]*string)                                                                 // 2. ［引数名］と、［その値が入る変数へのポインター］のマッピング（入れ物）を用意
	pArgsMap["p"] = fs1.String("p", "", "Program name. It is the file name under the 📁exercise folder.") // 3. ［引数名］を登録し、後でその値が入る変数へのポインターを取得
	pArgsMap["f"] = fs1.String("f", "", "Target file path.")

	subsequentTokens := tokens[1:] // 4. コマンドラインから先頭のコマンド名を取り除いた、［２つ目以降の単語の配列］を取得
	fs1.Parse(subsequentTokens)    // 5. ［２つ目以降の単語の配列］を、コマンドライン引数として解釈

	return tokens[0], pArgsMap
}
