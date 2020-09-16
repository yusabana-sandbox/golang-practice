package main

import (
	"github.com/yusabana-sandbox/golang-practice/had-experience-of-programming/chapter2"
	"github.com/yusabana-sandbox/golang-practice/had-experience-of-programming/chapter3"
	"github.com/yusabana-sandbox/golang-practice/had-experience-of-programming/chapter4"
)

func main() {
	chapter2.Do()
	chapter3.Do()
	chapter4.Do()
	// serverとして起動するので適宜利用するときだけOnにする
	//chapter4.DoNetHttp()
}
