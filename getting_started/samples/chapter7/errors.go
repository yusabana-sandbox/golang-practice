package chapter7

import (
	"fmt"
)

func DoSamples() {

	// panicとrecover 即時関数をdeferで呼び出した中で recover() 関数を実行してpanicじの指定したオブジェクトを取得する
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}() // <= 即時関数を実行

	panic("ERROR") // 文字列じゃなくても良い
}

type escape struct{}

func f() { g() }
func g() { panic(escape{}) }

func Hoge() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(escape); ok {
				println("Escaped")
			} else {
				panic(r)
			}
		}
	}()
	f()
}
