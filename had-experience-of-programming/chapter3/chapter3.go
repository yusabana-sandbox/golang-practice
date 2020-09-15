package chapter3

import (
	"fmt"
)

func Do() {
	fmt.Println("chapter3始まるよー")
	// 型のチェック(ポインタ等)
	checks()
	checksType()

	// 構造体
	doStructure()
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
	id     int
	detail string
	done   bool
}

func doStructure() {
	var task1 Task = Task{
		id:     1,
		detail: "ほげふが",
		done:   false,
	}
	fmt.Println(task1)

	var task2 Task = Task{1, "AAA", false}
	fmt.Println(task2)

	// 構造体を初期化しない場合はゼロ値で初期化される
	task3 := Task{}
	fmt.Println(task3.id, task3.detail, task3.done)

	// 構造体の前に&をつけてポインタ形にできる. ポインタと値を用途に応じて分ける
	var taskItself Task = Task{}   // Task型
	var taskPointa *Task = &Task{} // Taskのポインタ型
	var taskNew *Task = new(Task) // 組み込み関数new 構造体フィールドをゼロ値初期化しそのポインタを返す
	fmt.Println(
		taskItself,
		taskPointa,
		*taskPointa,
		taskNew,
	) // => {0  false} &{0  false} {0  false} &{0  false}

	// Task型の場合, 関数に値をコピーして渡されるためもとのtask11は変わらない
	fmt.Println("taskItself before:", taskItself.done)
	var doTask11 func(task Task) = func(task Task) {
		task.done = true
	}
	doTask11(taskItself)
	fmt.Println("taskItself after: ", taskItself.done) // falseのまま

	// Taskのポインタ型の場合
	fmt.Println("taskPointa before:", taskPointa.done)
	var doTask22 func(task *Task) = func(task *Task) {
		task.done = true
	}
	doTask22(taskPointa)
	fmt.Println("taskPointa after: ", taskPointa.done) // trueに変わる

	// コンストラクタ
	task := NewTask(3, "aaaaaa")
	fmt.Printf("%+v\n", task)
	// メソッド
	fmt.Printf("%s\n", task)

	fmt.Println(task) // &{3 aaaaaa false}
	task.Finish()     // Finish()はレシーバーのポインタに対する操作なので呼び出し側recieverも変わる
	fmt.Println(task) // &{3 aaaaaa true}
}

// コンストラクタ
func NewTask(id int, detail string) *Task {
	fmt.Println("コンストラクタとしてのNewTaskを実行")

	task := &Task{
		id: id,
		detail: detail,
		done: false,
	}

	return task
}

// メソッド
//値をレシーバーに設定しているのでメソッド内部でレシーバーの中身を変更しても呼び出し側の構造体には反映されない
// String()メソッドは特殊でfmt.Printlnとかで使われる. JavaでいうtoString()と同じイメージ
func (task Task) String() string {
	str := fmt.Sprintf("%d) %s, DONE: %v", task.id, task.detail, task.done)
	return str
}

// レシーバーをポインタ
// ポインタなのでレシーバーの変更を行うと、呼び出し側の構造体もそのまま変更される
func (task *Task) Finish() {
	task.done = true
}
