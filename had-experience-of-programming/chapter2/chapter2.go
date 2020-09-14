package chapter2

import (
	"errors"
	"fmt"
	"log"
)

func Do() {
	fmt.Println("chapter2始まるよー")

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

	// 関数リテラル(即時関数)
	func(i, j int) { fmt.Println(i + j) }(3, 5)
	// 関数リテラルの変数代入(Rubyのlambdaのようなものsumという名前で作る)
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

	// マップ
	mapFunc()

	// ポインタ
	n := 3
	fmt.Println("noPointa実行前, pointa実行前", n)
	noPointa(n) // この場合は値を渡しているので値がコピーされて渡される
	fmt.Println("noPointa実行後, pointa実行前", n)
	pointa(&n) // &をつけて呼び出しにアドレスを渡す
	fmt.Println("noPointa実行後, pointa実行後", n) // ポインタの参照を書き換えているのでnの値も変わる

	// defer
	defer fmt.Println("Deffer1で最後の処理")
	defer fmt.Println("Deffer2で最後の処理")
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

// マップ
func mapFunc() {
	var month map[int]string = map[int]string{}
	month[1] = "January"
	month[2] = "February"

	fmt.Println(month)

	// 2つ目の戻り値は指定したキーがマップに格納されているかをboolで返す
	m, ok := month[2]
	if ok {
		fmt.Println("monthはOKです", m)
	}

	// rangeを使うとfor文でkey, valueを取れる
	for key, value := range month {
		fmt.Printf("%d -> %s\n", key, value)
	}

	// キーが1のデータを削除
	delete(month, 1)
	fmt.Println(month)
}

// ポインタのチェック
func noPointa(num int) {
	num = 50
	fmt.Println("noPointa関数内:", num)
}
func pointa(num *int) {
	*num = 50 // ポインタが指している変数に値をセットする
	fmt.Println("pointa関数内:", *num) // ポインタが指している変数を表示
}
