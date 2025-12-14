package exercise

// GetOption - エンジンオプションの値を取得する
func GetOption(name string, engineOptions map[string]string) {
	value, exists := engineOptions[name]
	if exists {
		println(value)
	} else {
		println("Option not found.")
	}
}
