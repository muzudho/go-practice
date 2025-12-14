package exercise

// SetOption - オプションを設定する練習
func SetOption(name string, value string, engineOptions map[string]string) {
	//fmt.Printf("呼び出し前: engineOptions = %v\n", engineOptions)
	engineOptions[name] = value
	//fmt.Printf("呼び出し後: engineOptions = %v\n", engineOptions)
}
