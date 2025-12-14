package exercise

// EchoProxy - 外部プロセスの標準入出力をプロキシする練習
func SetOption(name string, value string, engineOptions map[string]string) {
	engineOptions[name] = value
}
