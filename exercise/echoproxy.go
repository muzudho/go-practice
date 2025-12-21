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

	externalProcess := exec.Command(externalProcessPath, parameters...) // 外部プロセスコマンド作成
	// ワーキング・ディレクトリーは特に指定なし

	exStdin, err := externalProcess.StdinPipe() // 外部プロセス標準入力パイプ取得
	if err != nil {
		panic(err)
	}
	defer exStdin.Close() // stdin を使い終わったら、外部プロセス標準入力パイプクローズ

	exStdout, err := externalProcess.StdoutPipe() // 外部プロセス標準出力パイプ取得
	if err != nil {
		panic(err)
	}
	defer exStdout.Close() // stdout を使い終わったら、外部プロセス標準出力パイプクローズ

	err = externalProcess.Start() // 外部プロセス起動
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}

	// FIXME: ゴルーチンを使っているが、終了処理が適切に行われていない。
	go receiveStdout(exStdout) // 外部プロセスの標準出力受信開始

	// Go言語では標準出力のUTF-8に対応していますが、VSCodeのターミナルはUTF-8に対応していないようです。
	// `chcp 65001`
	// そのため、外部プロセスの標準出力を受信しても、正しく表示されない場合があります。
	// その場合は、WindowsのコマンドプロンプトやPowerShellなど、UTF-8に対応したターミナルで実行してください。

	// FIXME: ゴルーチンを使っているが、終了処理が適切に行われていない。
	go receiveStdin(exStdin) // 外部プロセスの標準入力送信開始

	fmt.Print("外部プロセスと接続しました。文字を入力してください。\n")
	externalProcess.Wait()

    // きちんと閉じること。閉じないと呼出し元のプロセスを邪魔するかも。
	// Explicitly close pipes to trigger errors in goroutines
	exStdin.Close()
	exStdout.Close()

	fmt.Print("外部プロセスが終了しました。\n")

	// FIXME: 元のプロセスに戻ると、標準入力と標準出力は元に戻っているはずだが、入力ができない場合がある。
}

// receiveStdin - 標準入力受信
// `epStdin` - External process stdin
func receiveStdin(epStdin io.WriteCloser) {
    // 念入りに閉じること。
	defer epStdin.Close() // Ensure close if goroutine exits early
    
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を読み取るスキャナ作成

	// FIXME: scanner.Err() エラー処理が必要？
	for scanner.Scan() {
		command := scanner.Text() // １行読み取り。UTF-8文字列。
		
		// この書き方は１行で書ける。
		//epStdin.Write([]byte(command))
		//epStdin.Write([]byte("\n"))
		// こう書く
		_, err := epStdin.Write([]byte(command + "\n"))
		
		// スキャンしたら、ループから抜けることが必要。
		if err != nil {
			// Pipe closed/broken: external process ended, stop consuming input
			return
		}		
	}
	
	// エラーを処理する。
	if err := scanner.Err(); err != nil {
		// Handle scanner errors (e.g., I/O issues)
		fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
	}
}

// receiveStdout - 標準出力受信
// `epStdout` - External process stdout
func receiveStdout(epStdout io.ReadCloser) {
	defer epStdout.Close() // 念入りに閉じること。
    
    // UTF-8 文字は1バイト以上になることがあるから、１バイトずつ読込むのはよくない。
    // マルチバイト・バッファーを使う。
	const bufferSize = 1024
	buffer := make([]byte, bufferSize)
    
	//// FIXME: バッファサイズを1バイトにしているが、UTF-8文字列は1バイト以上になる場合がある。
	//var buffer [1]byte // これが満たされるまで待つ。1バイト。
	//p := buffer[:]

	for {
		n, err := epStdout.Read(buffer)
		//n, err := epStdout.Read(p)

		if err != nil {
			if err == io.EOF {
				return
			}
		  
			//if fmt.Sprintf("%s", err) == "EOF" {
			//	// End of file
			//	return
			//}

			panic(err)
		}

		if 0 < n {
		    fmt.Print(string(buffer[:n]))
			// 思考エンジンから１文字送られてくるたび、表示。
			//bytes := p[:n]
			//print(string(bytes))
		}
	}
	
	// ここには来ない。
}
