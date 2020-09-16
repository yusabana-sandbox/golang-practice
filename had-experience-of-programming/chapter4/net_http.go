package chapter4

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type WebPerson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func DoNetHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!!!!!")
	})
	http.HandleFunc("/persons", PersonHandler)

	http.ListenAndServe(":8000", nil)
}

// テンプレートのコンパイル
// ParseFiles()は戻り地にerrを返すが、
// t, err := template.ParseFiles("index.html")
// Must()を一緒に用いるとエラー位に戻り地じゃなくてパニックを発生させることができる
// そもそもテンプレートがないとかコンパイルできないとかは実行時エラーと言うよりはおかしい状態なのでpanicにするのがいい
// 一度コンパイルが通ることを確認したテンプレートであれば毎回エラー処理を書く必要性は低いのでMust()を使うのがいい
var t = template.Must(template.ParseFiles("index.html"))

/*
以下のコマンドを実行してPOSTするとファイルが生成される
POST
$ curl -vvvv http://localhost:8000/persons -d '{"id":1, "name": "ほげふが"}'

以下のコマンドはGETしてファイルに書き込まれた情報を取得する
$ curl -vvvv 'http://localhost:8000/persons/?id=1'
 */
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	// 処理の最後にBodyを閉じる
	defer r.Body.Close()

	fmt.Println("aaaa")

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person WebPerson
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.Id)
		// ioutilで一行で済む
		ioutil.WriteFile(filename, []byte(person.Name), 0666)
		//file, err := os.Create(filename)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file.Close()
		//
		//_, err = file.WriteString(person.Name)
		//if err != nil {
		//	log.Fatal(err)
		//}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Created!!")
	} else if r.Method == "GET" {
		// パラメータを取得(Atoi()を使ってstringからintに変換する)
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		person := Person{
			Id:   id,
			Name: string(b),
		}

		// templateに対して実行する。それをHTTPのレスポンスとして返す
		t.Execute(w, person)
	}
}
