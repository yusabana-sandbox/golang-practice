package samples

import "fmt"

func TestFunction() {
	fmt.Println("TestFunction Start ..............")

	rex, rey := func1(2, 3)
	fmt.Println(rex, rey)

	fmt.Println(namedFunc(5, 10))

	// 無名関数
	msg := "ほげ"
	func() {
		println("メッセージは " + msg)
	}()

	// 関数はファーストクラスオブジェクト
	fs := make([]func() string, 2) // stringを返す関数のスライス
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }
	for _, f := range fs {
		fmt.Println(f())
	}

	// 複数戻り値で値のコピーを利用
	var n, m = swap(10, 20)
	println(n, m)

	// 複数戻り値でポインタを利用
	nn, mm := 10, 20
	swap2(&nn, &mm)
	println(nn, mm) // ポインタを入れ替えたので値も変わる

	fmt.Println("TestFunction Stop ..............")
}

// ポインタを受け取る
func swap2(x *int, y *int) {
	*x, *y = *y, *x
}

// 値を入れ替える
func swap(x int, y int) (n int, m int) {
	n, m = y, x
	return
}

func func1(x int, y int) (int, int) {
	return x * 2, y * 2
}

// 名前付き戻り値
func namedFunc(x int, y int) (x2 int, y2 int) {
	x2, y2 = x*2, y*2
	return
}
