package chapter2

import (
	"errors"
	"fmt"
	"log"
)

func Do() {

	fmt.Println("HOGEですよねー")

	a, b, _ := multipleReturn()
	fmt.Println(a, b)

	div1, err := div(4, 2)
	//div1, err := div(4, 0) // errorになる
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(div1)

	div2, err2 := divNamed(6, 2)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(div2)

	// 関数リテラル
	func(i, j int) { fmt.Println(i + j) }(3,5)
	// 関数リテラルの変数代入
	var sum func(i, j int) = func(i, j int) {
		fmt.Println(i + j)
	}
	sum(2, 4)
}

func multipleReturn() (string, int, int) {
	return "hoge", 3, 4
}

// エラーの発生のサンプル
func div(i, j int) (int, error) {
	if j == 0 {
		// 自作のエラーを返す
		return 0, errors.New("divided by zero")
	}

	return i / j, nil
}

// 名前付き戻り値
func divNamed(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divided by zero of divNamed")
		return
	}
	result = i / j
	return
}
