package main

import (
	"flag"
	"fmt"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter4"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter5"
	"os"
	"strings"
)

// 設定される変数のポインタを取得
//var msg string // 1 こういうとり方も可能
var msg = flag.String("msg", "デフォルト", "説明")
var n int

func init() {
	// ポインタを指定して設定を予約
	// flag.StringVar(&msg, "msg", "デフォルト", "説明") // 1 こういうとり方も可能
	flag.IntVar(&n, "n", 5, "回数")

	fmt.Println("INIT1")
}

func main() {
	fmt.Println(os.Args)

	// parseで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))

	// flag.Parse()しないとflag.Args()でとれない
	// https://qiita.com/178inaba/items/7f412a1acb435a202f5c
	fmt.Println(flag.Args())

	fmt.Println("CHAPTER 444444444")
	chapter4.DoImports()

	fmt.Println("CHAPTER 5555555555")
	chapter5.DoArgs()
}
