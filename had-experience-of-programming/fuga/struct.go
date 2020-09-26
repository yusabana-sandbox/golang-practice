package fuga

import "fmt"

type Calc struct {
	a, b int
}

// 型を埋め込む(マージする)
type MyStruct struct {
	Calc
}

type MyInt int

// Calc構造体に紐づく関数(メソッド)
func (p Calc) Add() int {
	return p.a + p.b
}

// MyInt型に紐づく関数(メソッド)
func (m MyInt) Add(n int) MyInt {
	return m + MyInt(n)
}

func RunStruct() {
	fmt.Println("=== RunStruct ===")
	p := Calc{a: 3, b: 10}
	var m MyInt = 1

	fmt.Println(p.Add()) // => 13
	fmt.Println(m.Add(33)) // => 34

	var s MyStruct
	s.a = 5
	s.b = 4
	fmt.Println(s.Add()) // => 9
}
