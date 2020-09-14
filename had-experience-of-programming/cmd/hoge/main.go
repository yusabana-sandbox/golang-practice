package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("HOGEですよねー")

	a, b, _ := multipleReturn()
	fmt.Println(a, b)

	div1, err := div(4, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(div1)
}


func multipleReturn() (string, int, int) {
	return "hoge", 3, 4
}


func div(i float32, j int) (int, error) {
	if j == 0 {
		// 自作のエラーを返す
		return 0, errors.New("divided by zero")
	}

	return int(i) / j, nil
}
