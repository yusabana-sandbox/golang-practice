package main

import "fmt"

func main() {
	str := "hello, world!!あ"
	vec := []byte(str)
	fmt.Println(str)
	fmt.Println(vec)

	DifferentMakeNew()

	Map()

}

func Map() {
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
