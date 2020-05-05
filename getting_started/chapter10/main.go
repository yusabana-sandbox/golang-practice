package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type MyHandler struct{
	http.Handler
}
func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO, HANDLERで定義")
}


func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO, HANDLER_FUNCで定義")
}

func handlerOmikuji(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(6) +1

	switch num {
	case 6:
		fmt.Fprint(w, "おみくじ結果は大吉")
	case 5,4:
		fmt.Fprint(w, "おみくじ結果は中吉")
	case 3,2:
		fmt.Fprint(w, "おみくじ結果は吉")
	case 1:
		fmt.Fprint(w, "おみくじ結果は凶")
	}
}

func init() {
	// 乱数の初期化
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/", handlerFunc)

	handler := MyHandler{}
	http.Handle("/2", handler)

	// おみくじ
	http.HandleFunc("/omikuji", handlerOmikuji)

	http.ListenAndServe(":8080", nil)
}



