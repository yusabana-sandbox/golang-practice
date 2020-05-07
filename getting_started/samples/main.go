package main

import (
	"flag"
	"fmt"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter11"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter4"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter5"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter6"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter7"
	"github.com/yusabana-sandbox/golang-practice/getting_started/samples/chapter9"
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

	fmt.Println("CHAPTER 6666666666")
	chapter6.DoSamples()

	fmt.Println("CHAPTER 7777777777")
	chapter7.DoSamples()

	fmt.Println("CHAPTER 99999999999")
	// promptは実行が止まってしまうので一旦無効に
	//chapter9.DoPrompt()
	chapter9.DoLock()

	fmt.Println("CHAPTER 11")
	chapter11.DoSamples()
}
