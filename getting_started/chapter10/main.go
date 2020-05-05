package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type MyHandler struct {
	http.Handler
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO, HANDLERで定義")
}

var tmpl = template.Must(template.New("sign").
	Parse("<html><body>{{.Name}}さんの運勢は「<strong>{{.Omikuji}}</strong>」です</body></html>"))

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("p")

	data := struct {
		Name    string
		Omikuji string
	}{
		Name:    name,
		Omikuji: omikuji(),
	}
	tmpl.Execute(w, data)
}

func handlerOmikuji(w http.ResponseWriter, r *http.Request) {
	result := omikuji()
	fmt.Fprint(w, "おみくじ結果は"+result)
}

func omikuji() string {
	// rand 0-5の6パターンなので+1してswitchでサイコロ飲めのようにおみくじを決める
	num := rand.Intn(6) + 1
	var r string
	switch num {
	case 6:
		r = "大吉"
	case 5, 4:
		r = "中吉"
	case 3, 2:
		r = "吉"
	case 1:
		r = "凶"
	}
	return r
}

func handleMyJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Name string `json:"N"` // 個々のタグで指定したキーでjsonに変換される
		Msg  string `json:"msg"`
	}{
		Name: "Hoge",
		Msg:  "hello",
	}

	// encodeするとレスポンスがJSONになる?!
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}

func main() {
	// 乱数の初期化
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handlerFunc)

	handler := MyHandler{}
	http.Handle("/2", handler)

	// おみくじ
	http.HandleFunc("/omikuji", handlerOmikuji)

	// JSONサンプル
	http.HandleFunc("/myjson", handleMyJson)

	http.HandleFunc("/handle_test", handleTest)
	http.ListenAndServe(":8080", nil)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, net/http!!")
}
