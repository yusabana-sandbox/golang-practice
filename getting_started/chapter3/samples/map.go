package samples

import "fmt"

func TestMap() {
	fmt.Println("TestMap Start ..............")

	m := map[string]int{"x": 10, "y": 20}
	fmt.Println(m)
	fmt.Println(m["x"])

	m["z"] = 300
	n, ok := m["z"]
	fmt.Println(n, ok) // 300 true

	delete(m, "z")
	n, ok = m["z"]
	fmt.Println(n, ok) // 0 false   intのvalueなのでゼロ値が0となる

	// 二次元配列
	nums := [][]int{{1, 2, 3}, {2, 3}}
	// mapの値がslice
	deeps := map[string][]int{"a": {1, 2, 3}, "b": {11, 22, 33}}
	fmt.Println(nums)
	fmt.Println(deeps)

	fmt.Println("TestMap Stop ..............")
}
