package chapter5

import "fmt"

func DoArgs() {
	msg := "!!!"
	defer fmt.Println(msg)
	msg = "World"
	defer fmt.Println(msg)
	fmt.Println("hello")

}
