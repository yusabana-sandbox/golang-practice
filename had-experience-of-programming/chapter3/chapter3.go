package chapter3

import "fmt"

func Do() {
	fmt.Println("chapter3始まるよー")
	checks()
	checksType()
}

// 型がintだからどちらでも通るというパターン
func checks() {
	id := 1
	priority := 3
	var checks func(id, priority int) = func(id, priority int) {
		fmt.Println(id + priority)
	}

	// 逆にしても通っちゃう
	checks(id, priority) // ok
	checks(priority, id) // ok

	// 構造体
	doStructure()
}

type Id int
type Priority int

func checksType() {
	var checks func(id Id, priority Priority) = func(id Id, priority Priority) {
		// 型をキャストしないと計算できない（型が違うので）
		fmt.Println(int(id) + int(priority))
	}

	var id Id = 1
	var priority Priority = 5

	checks(id, priority)
	//checks(priority, id) // コンパイルエラー
}

// 構造体関連
type Task struct {
	id int
	detail string
	done bool
}
func doStructure() {
	var task1 Task = Task{
		id: 1,
		detail: "ほげふが",
		done: false,
	}
	fmt.Println(task1)

	var task2 Task = Task{1, "AAA", false}
	fmt.Println(task2)

	// 構造体を初期化しない場合はゼロ値で初期化される
	task3 := Task{}
	fmt.Println(task3.id, task3.detail, task3.done)
}

