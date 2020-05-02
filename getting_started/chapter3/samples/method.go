package samples

import "fmt"

type T int

func (t *T) f() { println("hi") }
func (t *T) f2() {
	*t = *t + 10
}

type MyInt int

func (n *MyInt) inc() {
	*n++
}

// Hex型を定義して
type Hex int

// Hexの型に対してString()メソッドを定義する
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}
func TestMethod() {
	fmt.Println("TestMethod Start ..............")

	var hex Hex = 100
	fmt.Println(hex.String())

	var v T // ゼロ値で 0 が初期化

	println(v) // 0

	v.f()    // hi
	(&v).f() // hi

	v.f2()     // vの値が10に変更される
	println(v) // 10

	fmt.Println("-------")
	var n MyInt
	println(n)
	n.inc()
	println(n)

	// メソッド値
	f1 := hex.String // ここでhexを束縛してメソッドを値とする
	fmt.Println("f1() => ", f1())

	// メソッド式
	f2 := Hex.String // Hex型のStringメソッドを表す式（レシーバを第一引数とした関数）
	fmt.Printf("f2() =>\n  %T\n  %s\n", f2, f2(hex))

	fmt.Println("TestMethod Stop ..............")
}
