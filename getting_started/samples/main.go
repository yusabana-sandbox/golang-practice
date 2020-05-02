package main

import (
	"fmt"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter4"
)

func init() {
	fmt.Println("INIT1")
}

func init() {
	fmt.Println("INIT2")
}

func main() {
	chapter4.DoImports()
}
