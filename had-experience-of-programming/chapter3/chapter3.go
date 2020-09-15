package chapter3

import (
	"fmt"
	"io"
)

func Do() {
	fmt.Println("chapter3始まるよー")
	// 型のチェック(ポインタ等)
	checks()
	checksType()

	// 構造体
	doStructure()

	// Cast
	doCast()

	// Type Assertion 型の検査
	doTypeAssertion("aiueo")
	doTypeAssertion(1234)
	doTypeAssertion(false)
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

type User struct {
	firstName string
	lastName  string
}
type Task2 struct {
	id     int
	detail string
	done   bool
	//User   *User // Userを定義する場合はこのようにする.埋め込みじゃなくてプロパティとして設定
	*User // *Userとすることで埋め込む Task2にUserが埋め込まれる
}

type Stringer interface {
	String() string
}

// インターフェースの埋め込み（型定義だけ）
// ioのReaderとWriterを持った型
type ReadWriter interface {
	io.Reader
	io.Writer
}

// コンストラクタ
func NewTask(id int, detail string) *Task {
	fmt.Println("コンストラクタとしてのNewTaskを実行")

	task := &Task{
		id:     id,
		detail: detail,
		done:   false,
	}

	return task
}

func NewUser(firstName, lastName string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
	}
}

func NewTask2(id int, detail, firstName, lastName string) *Task2 {
	task := &Task2{
		id:     id,
		detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
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

func (u *User) FullName() string {
	fullName := fmt.Sprintf("%s %s", u.firstName, u.lastName)
	return fullName
}

// インターフェース
func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}

func doStructure() {
	var task1 Task = Task{
		id:     1,
		detail: "ほげふが",
		done:   false,
	}
	fmt.Println(task1)

	var task12 Task = Task{1, "AAA", false}
	fmt.Println(task12)

	// 構造体を初期化しない場合はゼロ値で初期化される
	task3 := Task{}
	fmt.Println(task3.id, task3.detail, task3.done)

	// 構造体の前に&をつけてポインタ形にできる. ポインタと値を用途に応じて分ける
	var taskItself Task = Task{}   // Task型
	var taskPointa *Task = &Task{} // Taskのポインタ型
	var taskNew *Task = new(Task)  // 組み込み関数new 構造体フィールドをゼロ値初期化しそのポインタを返す
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

	// インターフェース
	// taskはString()メソッドを定義しているのでPrintに渡すことができる
	// Stringerインターフェースを実装していることになる
	Print(task)

	// Structure(構造体)の埋め込み
	task2 := NewTask2(1, "hoho", "first_hoge", "last_fuga")
	fmt.Println(task2.firstName)  // 構造体に埋め込まれるのでそのまま呼べる
	fmt.Println(task2.lastName)   // 構造体に埋め込まれるのでそのまま呼べる
	fmt.Println(task2.FullName()) // 構造体に埋め込まれるのでそのまま呼べる
	fmt.Println(task2.User)       // Task構造体から埋め込まれたUserにもアクセスできる
}

func doCast() {
	var i uint8 = 3
	var j uint32 = uint32(i)
	fmt.Println(j) // => 3

	var s string = "abc"
	var b []byte = []byte(s)
	fmt.Println(b) // => [97 98 99]

	// Cannot convert expression of type 'string' to type 'int'
	// cast出来ないときはpanicが発生する
	//var i int = int(s)
}

// Type Assertion
// レシーバー.(string) のような感じで型を検査する
// 第1戻り値は判定が成功した場合その方に変換された値が返る
// 第2戻り値は判定が成功したかどうかの真偽値が返る
//   第2戻り値を取らなかったら判定失敗時はpanicとなる
func doTypeAssertion(value interface{}) {
	s, ok := value.(string)

	// 第2戻り値を取っているのでassertionに失敗してもpanicは起こらず条件判定できる
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string: %v\n", value)
	}
}
