package main

import (
	"had-experience-of-programming/chapter2"
	"had-experience-of-programming/chapter3"
	"had-experience-of-programming/chapter4"
)

func main() {
	chapter2.Do()
	chapter3.Do()
	chapter4.Do()
	// serverとして起動するので適宜利用するときだけOnにする
	chapter4.DoNetHttp()
}
