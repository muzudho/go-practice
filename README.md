# go-practice

Go言語の練習（＾～＾）

```shell
go build

## Run
##      -p はプログラム名。📁 exercise 下のファイル名が入る。
go-practice -p hello
    ## world

>>> quit
```

```shell
go-practice

>>> hello
    ## world

>>> quit
```

```shell
go-practice -p set-option -n engine -v banana

[banana] get-option -n engine
    ## banana

[banana] quit
```

```shell
## go-practice -p fmt
## go-practice -p mapping-char

## >>> character -s "B"
## >>> fmt -s "Hello, world!!"
## >>> fmt -s "日本語だったらどうなる（＾ｑ＾）？"
## >>> fmt -s "apple バナナ Cherry"
## >>> fmt -s "体重は90kgぐらいかだぜ（＾ｑ＾）？"

```shell
go-practice -p set-option -n engine -v banana

## 例えば、実行ファイルのパスは 📄 `Z:/muzudho-github.com/muzudho/go-practice/go-practice.exe` だとします。  
[banana] echo-proxy -f Z:/muzudho-github.com/muzudho/go-practice/go-practice.exe

>>> hello
    ## world

>>> quit

[banana] 
```
```


## 外部リンク

* 🌏 [Go言語 - 日本語文字列の操作](https://hake.hatenablog.com/entry/20150826/p1)
* 🌏 [Golang | How to find the index of rune in the string?](https://www.geeksforgeeks.org/golang-how-to-find-the-index-of-rune-in-the-string/)
* 🌏 [Check If the Rune is a Letter or not in Golang](https://www.geeksforgeeks.org/check-if-the-rune-is-a-letter-or-not-in-golang/)
