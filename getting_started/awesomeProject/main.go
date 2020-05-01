package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	helloWorld()

	// bit演算をiotaで定義するみたいに式の途中にも使える
	const (
		a = 1 << iota // a = 0 << 1
		b             // b = 1 << 2
		c             // c = 2 << 4
		d             // d = 3 << 8
	)
	// 1 2 4 8と表示される
	fmt.Println(a, b, c, d)

	switchCase(1)
	switchCase(2)
	switchCase(3)

	showOddEven()

	rand.Seed(time.Now().UnixNano())
	showOmikuji(rand.Intn(6) + 1) // 0-5までの乱数が出るので+1する
}

func switchCase(aa int) {
	// こんな感じでcaseに式を書ける
	switch {
	case aa == 1, aa == 2:
		fmt.Println("OK")
	default:
		fmt.Println("DEFAULT")
	}
}

func showOddEven() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - 偶数\n", i)
		} else {
			fmt.Printf("%d - 奇数\n", i)
		}
	}
}

func showOmikuji(num int) {
	fmt.Printf("%d = ", num)
	switch num {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("吉")
	case 1:
		fmt.Println("凶")
	}
}
