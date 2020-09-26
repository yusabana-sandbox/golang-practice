package fuga

import "fmt"

type Human interface {
	hello()
	walk()
}

type Masaru struct {
	name string
	age int
}

func (m Masaru) hello() {
	fmt.Printf("%s%dです\n", m.name, m.age)
}

func (m Masaru) walk() {
	fmt.Printf("%sがトコトコ...\n", m.name)
}

func (m Masaru) shout() {
	fmt.Println("うぉーーーー")
}

func RunInterface() {
	fmt.Println("=== RunInterface ===")
	// 構造体をインターフェースに設定する
	m := Masaru{"まさる", 20}
	var h Human = m

	m.hello()
	m.walk()
	m.shout()
	fmt.Println(m.name)

	fmt.Println("------")

	h.hello()
	h.walk()
	// 以下はHumanインターフェースにないので使えない
	//h.shout()
	//fmt.Println(h.name)


}
