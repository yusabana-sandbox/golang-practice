package main

import "fmt"

func main() {
	fmt.Println("HOGEですよねー")

	a, b, _ := multipleReturn()

	fmt.Println(a, b)
}


func multipleReturn() (string, int, int) {
	return "hoge", 3, 4
}
