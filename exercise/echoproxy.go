package exercise

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// EchoProxy - 外部プロセスの標準入出力をプロキシする練習
func EchoProxy(externalProcessPath string) {
	// echo-proxy コマンドは、パラメーターは受け取っていません。
	parameters := []string{} // No parameters, as per comment

	// 外部プロセスコマンド作成
	// ワーキング・ディレクトリーは特に指定なし
	externalProcess := exec.Command(externalProcessPath, parameters...)

	// 外部プロセス標準入力パイプ取得
	exStdin, err := externalProcess.StdinPipe()
	if err != nil {
		panic(err)
	}
	// stdin を使い終わったら、外部プロセス標準入力パイプクローズ
	defer exStdin.Close()

	// 外部プロセス標準出力パイプ取得
	exStdout, err := externalProcess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	// stdout を使い終わったら、外部プロセス標準出力パイプクローズ
	defer exStdout.Close()

	// 外部プロセス起動
	err = externalProcess.Start()
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}

	// doneチャネルでゴルーチンを制御
	done := make(chan struct{})

	// ゴルーチンを使って、外部プロセスの標準出力受信開始
	go receiveStdout(exStdout, done)

	// Go言語では標準出力のUTF-8に対応していますが、VSCodeのターミナルはUTF-8に対応していないようです。
	// `chcp 65001`
	// そのため、外部プロセスの標準出力を受信しても、正しく表示されない場合があります。
	// その場合は、WindowsのコマンドプロンプトやPowerShellなど、UTF-8に対応したターミナルで実行してください。

	// ゴルーチンを使って、外部プロセスの標準入力送信開始
	go receiveStdin(exStdin, done)

	fmt.Print("外部プロセスと接続しました。文字を入力してください。\n")
	externalProcess.Wait()

	// プロセス終了をシグナルしてゴルーチンを停止
	close(done)

	// 明示的にパイプを閉じてエラーをトリガー
	exStdin.Close()
	exStdout.Close()

	// ガベージコレクションを強制実行して、os.Stdinの状態をクリーンにする
	//runtime.GC()

	fmt.Print("外部プロセスが終了しました。呼び出し元プロセスの標準入力がクリーンになるまで、改行を送ってきてください。\n")

	// それが終わると、os.Stdinの状態がクリーンになり、親プロセスで即入力可能
}

// receiveStdin - 標準入力受信
// `epStdin` - External process stdin
// `done` - 終了シグナルチャネル
func receiveStdin(epStdin io.WriteCloser, done <-chan struct{}) {
	// Ensure close if goroutine exits early
	defer epStdin.Close()

	// 標準入力を読み取るスキャナ作成
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-done:
			// プロセス終了シグナル受信: 残り入力をドレインして終了
			//fmt.Println("[ｅchoproxy.go > ｒeceiveStding > done] 標準入力を読み込みます")
			for scanner.Scan() {
				// 無視（バッファクリア）
			}
			return
		default:
			//fmt.Println("[ｅchoproxy.go > ｒeceiveStding > default] 標準入力を読み込みます")
			if !scanner.Scan() {
				// エラーを処理する。
				if err := scanner.Err(); err != nil {
					// Handle scanner errors (e.g., I/O issues)
					fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
				}

				break
			}

			// １行読み取り。UTF-8文字列。
			command := scanner.Text()
			_, err := epStdin.Write([]byte(command + "\n"))
			if err != nil {
				// パイプ閉鎖エラー: 早期終了
				return
			}
		}
	}
}

// receiveStdout - 標準出力受信
// `epStdout` - External process stdout
// `done` - 終了シグナルチャネル
func receiveStdout(epStdout io.ReadCloser, done <-chan struct{}) {
	defer epStdout.Close()

	// UTF-8 文字は1バイト以上になることがあるから、１バイトずつ読込むのはよくない。
	// マルチバイト・バッファーを使う。
	const bufferSize = 1024
	buffer := make([]byte, bufferSize)

	for {
		select {
		case <-done:
			return
		default:

			n, err := epStdout.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}

			if 0 < n {
				fmt.Print(string(buffer[:n]))
			}
		}
	}

	// ここには来ない。
}
