package samples

import "fmt"

func TestSlice() {
	fmt.Println("TestSlice Start ..............")

	var ns1 []int
	ns1 = make([]int, 3, 10)
	fmt.Println(ns1)

	// スライスリテラルで初期化
	// 配列との違いは型の指定で []int(10,20) となっていること、配列は [...]int(10,20)となっている
	ns2 := []int{10, 20, 30, 40, 50}
	ns3 := []int{5: 50, 10: 100}
	fmt.Println(ns2)
	fmt.Println(ns3)

	ns := []int{10, 20, 30, 40, 50}
	println(ns) // [5/5]0xAAAAAAAA [len/cap]となっている
	fmt.Println(ns, len(ns), cap(ns))

	ns = append(ns, 60, 70, 80)
	println(ns) // [8/10]0xAAAAAAA [len/cap]で表示される
	fmt.Println(ns, len(ns), cap(ns))

	// スライスへのスライス演算
	ns11 := []int{10, 20, 30, 40, 50}
	n, m := 2, 4
	fmt.Println(ns11[n:])  // [30 40 50]
	fmt.Println(ns11[:m])  // [10 20 30 40]
	ms11 := ns11[:n:m]     // ns[0:n:m] => ns[0:2:4] を指定した感じcapを4にする
	fmt.Println(ms11)      // [30 40]
	fmt.Println(cap(ms11)) // 4

	fmt.Println("TestSlice Stop ..............")
}
