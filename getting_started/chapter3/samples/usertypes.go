package samples

import (
	"fmt"
	"net/http"
	"time"
)

type Person struct {
	Name   string
	Number int
}

func TestUserTypes() {
	fmt.Println("TestUserTypes Start ..............")

	// 規定型とユーザー定義型の相互キャスト可能
	type MyInt int
	var n int = 100
	m := MyInt(n)
	n = int(m)

	// 10秒を表す(time.Duration型) 10のデフォルトの型からユーザー定義型へキャストできる場合はキャスト不要
	d := 10 * time.Second
	fmt.Println(d) // 10s
	println(d)     // 10000000000

	// 課題のやつ ユーザー定義型
	type Point = int
	type Points []Point

	type AppClient = http.Client
	fmt.Printf("%T\n", AppClient{}) // 型として http.Clientが表示される

	// ユーザー定義型で構造体を利用
	var p Person = Person{"aaa", 3}
	fmt.Println(p.Name)
	fmt.Println(p.Number)

	fmt.Println("TestUserTypes Stop ..............")
}
