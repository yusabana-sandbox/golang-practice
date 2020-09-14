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
	func(i, j int) { fmt.Println(i + j) }(3, 5)
	// 関数リテラルの変数代入
	var sum func(i, j int) = func(i, j int) {
		fmt.Println(i + j)
	}
	sum(2, 4)

	// 配列(スライスではない)
	arr()

	// スライス
	slice()

	// Range
	rng()

	// 値の切り出し
	// string, 配列, スライスから値を部分的に取り出せる
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("", s[2:4])      // [2 3]
	fmt.Println("", s[0:len(s)]) // [0 1 2 3 4 5]
	fmt.Println("", s[:3])       // [0 1 2]
	fmt.Println("", s[3:])       // [3 4 5]
	fmt.Println("", s[:])        // [0 1 2 3 4 5]

	// sumFuncは可変長引数
	fmt.Println(sumFunc(1, 2, 3, 4, 5))
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

// 配列
func arr() {
	fmt.Println("配列について")
	// どちらも同じ型
	arr1 := [4]string{"A", "B", "c", "d"}
	// この書き方は可変長のものではない(スライスではない)
	arr2 := [...]string{"A", "B", "c", "d"}

	sampleArr(arr1)
	sampleArr(arr2)
}

// 型引数に配列自体が渡されているのでコピーが渡されている
func sampleArr(arr [4]string) {
	fmt.Println(arr)
}

func slice() {
	fmt.Println("スライスについて")
	s := []string{"a", "b", "C", "D"}
	fmt.Println(s)
	s = append(s, "PIYO")
	fmt.Println(s)
}

func rng() {
	fmt.Println("レンジについて")
	arr1 := [4]string{"a", "b", "c", "d"}

	for i, s := range arr1 {
		// i 添字, s 文字列
		fmt.Println(i, s)
	}
}

// 可変長引数
func sumFunc(nums ...int) (result int) {
	// numsは var nums []int と同じ感じ
	for _, n := range nums {
		result += n
	}
	return
}
