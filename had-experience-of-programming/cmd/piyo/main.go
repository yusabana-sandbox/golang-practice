package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

func main() {
	//str := "hello, world!!あ"
	//vec := []byte(str)
	//fmt.Println(str)
	//fmt.Println(vec)
	//
	//DifferentMakeNew()
	//
	//DoMap()
	//
	//DoConst()
	//
	//DoInterface()

	DoHttp()
}

func DoHttp() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	// http://127.0.0.1:8000/counter でアクセス

	// ArgServerをhttp.HandlerFuncに型変換することでhttp.Handleに渡せるようになる
	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.ListenAndServe(":8000", nil)
}
// 単純なカウントサーバ
type Counter struct {
	n int
}
// レシーバは呼び出し側にも伝わるようにポインタである必要がある
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "カウント => %d\n", ctr.n)
}
// 以下のようにArgServerを定義することでHandlerFuncと同じシグネチャを持つようになったので
// ArgServerをHandlerFunc内のメソッドを使用するためにHandlerFunc型に変換可能になった
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for i, s := range os.Args {
		fmt.Fprintf(w, "%d => %s", i, s)
	}
}


func DoInterface() {
	hoge := Sequence{1, 10, 33, 100, 1000, 20}

	fmt.Println(hoge)
}
type Sequence []int
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Sequence) String() string {
	sort.Sort(s) // Len, Less, Swap を実装しているのでsortパッケージのソートで実行できる
	// このやり方だと簡易的にコードを軽量化できる
	// sort.IntSlice(s).Sort() // このようなソートにすることでLen, Less, Swapを実装しなくても良い

	//str := "["
	//for i, elem := range s {
	//	if i > 0 {
	//		str += " "
	//	}
	//	str += fmt.Sprint(elem)
	//}
	//return str + "]"
	//
	// 上記のコメントアウトしている処理はSequenceを純粋な[]intのスライスに変換してSprintすることで[]intのString()によって出力できる
	// 変換重要。変換しなければSprintはSequenceのStringメソッドを見つけ出し呼び出しを無限に行ってしまう無限ループに陥る
	// この2つの型(Sequenceと[]int)は名前を除けば同一であるため、これらの型の間における変換は問題なく行われます。この変換では新しい値が作られることはなく、既存の値が一時的に新しい型を持つかのような働きをします。(これと異なる変換もあります。たとえば整数を浮動小数点へ変換したときは新しく値が作成されます。)
	return fmt.Sprint([]int(s))
}

// 独自型のByteSizeの出力を単位付きで表示させるやつ
func DoConst() {
	var b0 ByteSize = 100
	var b1 ByteSize = 10000.0
	var b2 ByteSize = 1024
	var b3 ByteSize = 1048576
	var b4 ByteSize = 1 << 40
	fmt.Println(b0, b1, b2, b3, b4) // => 100.000B 9.766KB 1.000KB 1.000MB 1.000TB
}

type ByteSize float64

const (
	_           = iota // 一番目の値はブランク識別子に代入して無視
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// String()を上書きして出力形式を変換できる
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.3fYB", float64(b/YB))
	case b >= ZB:
		return fmt.Sprintf("%.3fZB", float64(b/ZB))
	case b >= EB:
		return fmt.Sprintf("%.3fEB", float64(b/EB))
	case b >= PB:
		return fmt.Sprintf("%.3fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%.3fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%.3fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%.3fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%.3fKB", float64(b/KB))
	}
	return fmt.Sprintf("%.3fB", float64(b))
}

func DoMap() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	var seconds int
	var ok bool

	tz := "EST"
	//複数代入式 ok には真偽値が入る
	seconds, ok = timeZone[tz]
	fmt.Println(timeZone, ok, seconds)
}

// make と newの違い
func DifferentMakeNew() {
	var p *[]int = new([]int)      // スライス構造の割り当て(*p == nil)。あまり使わない。
	var v []int = make([]int, 100) // スライスvは100個のintを持つ配列への参照
	fmt.Println(p, v)

	// 必要以上に複雑な書き方
	var p2 *[]int = new([]int)
	*p2 = make([]int, 100, 100)
	// p2がポインタで、*p2が値
	fmt.Println(p2, *p2)

	// 一般的な書き方
	v2 := make([]int, 100)
	fmt.Println(v2)
}

// サンプルのCompare関数
// Compare は2つのバイト配列を辞書的に比較して整数を返します。
// 結果は、a == bのとき0、a < bのとき-1、a > bのとき+1
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
