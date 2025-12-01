@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

move ファイルを移動中だぜ（＾～＾）...
move "../go-practice.exe" "C:/MuzudhoWorks/go-practice"
if %errorlevel%==0 (
    echo 移動完了！ よしよし（＾▽＾）
) else (
    echo エラー出たぜ... 確認してな。
)
