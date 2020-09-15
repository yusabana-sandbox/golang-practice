package chapter4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// 以下のコマンドで実行する
// $ curl -vvvv http://localhost:8000/persons -d '{"id":1, "name": "ほげふが"}'
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	// 処理の最後にBodyを閉じる
	defer r.Body.Close()

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
	}
}
