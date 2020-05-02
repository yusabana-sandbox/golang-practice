package main

import (
	"fmt"
	"github.com/yusabana-sandbox/golang-practice/getting_started/chapter3/samples"
)

func main() {
	sample()

	fmt.Println("==================")

	samples.TestSlice()
	samples.TestMap()
	samples.TestUserTypes()
	samples.TestFunction()
	samples.TestMethod()
}

func sample() {
	var a, b, c bool
	a = true
	b = false
	c = true
	condition := a && b || !c

	println(condition)

	var p struct {
		name string
		age  int
	}
	p.name = "AAA"
	p.age = 13

	// ゼロ値で初期化
	var ns1 [5]int
	fmt.Println(ns1)

	// 配列リテラルで初期化
	ns2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(ns2)

	// 要素数から配列のサイズを推論
	ns3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(ns3)
	// 5番目が50、10番目が100でほかが0の要素数11の配列
	ns4 := [...]int{5: 50, 10: 100}
	fmt.Println(ns4)

	// 以下はスライス 配列の一部を切り出したデータ構造
	fmt.Println(ns2[1:3])
	println(ns2[1:3])
}
