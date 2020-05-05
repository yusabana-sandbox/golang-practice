package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func handleMyJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Name string `json:"N"` // 個々のタグで指定したキーでjsonに変換される
		Msg string `json:"msg"`
	}{
		Name: "Hoge",
		Msg: "hello",
	}

	// encodeするとレスポンスがJSONになる?!
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("p")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("%sさんの運勢は「大吉」です", val),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Erro:", err)
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

	// JSONサンプル
	http.HandleFunc("/myjson", handleMyJson)
	http.HandleFunc("/req", handleReq)


	http.ListenAndServe(":8080", nil)
}




