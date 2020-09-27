package fuga

import (
	"errors"
	"fmt"
	"os"
)

func RunErrors() {
	// 存在しないファイルを指定必ず以下のエラーに入るはず
	file, err := os.Open("__AAAAAAA.txt")
	// errorインターフェースにはError()というメソッドがある
	if err != nil {
		fmt.Println("エラーよ")
		fmt.Println(err.Error()) // open __AAAAAAA.txt: no such file or directory
		return
	}

	// ファイルのクローズ処理をdeferで
	defer file.Close()

	// エラーではないときの処理
	fmt.Println("ファイルが開いた")

	//独自のエラーを返す場合
	// エラーの場合 errorだけを返す場合はこのように簡潔に書ける
	if e := test(); e != nil {
		fmt.Println(e.Error())
	}
}

func test() error {
	return errors.New("errors.Newでエラーを生成する")
}

