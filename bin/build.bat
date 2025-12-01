@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo `C:/MuzudhoWorks/go-practice` に `go-practice.exe` ファイルを作るぜ（＾～＾）...
cd ..
go build -o C:/MuzudhoWorks/go-practice
cd ./bin
echo go build したぜ（＾～＾）！
pause
