package chapter4

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func Do() {
	fmt.Println("chapter4始まるよー")

	doMarshal()
	doUnMarshal()

	doFiles()
}

func doMarshal() {
	person := &Person{
		Id:      1,
		Name:    "Goper",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	// structureにタグ付けしているのでその使用で出力される
	fmt.Println(string(b)) // => {"id":1,"name":"Goper","age":5}
}

func doUnMarshal() {
	var person Person

	b := []byte(`{"id":1,"name":"Gopher","age":5}`)

	err := json.Unmarshal(b, &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person) // => {1 Gopher 5 }
}

func doFiles() {
	// ファイル生成
	file, err := os.Create("./file.txt")

	if err != nil {
		log.Fatal(err)
	}
	// プログラムが終わったらファイルを閉じる
	defer file.Close()

	// defer でプログラムが終わったらファイルを閉じる
	// 書き込むデータを[]byteで利用する
	message := []byte("hello world!!!\n")

	// Writeで書き込む
	_, err = file.Write(message)
	// WriteStringというのもある
	//_, err = file.WriteString("hello world!!\n")
	// fmt.Fprintで直接fileに書き込める
	//_, err = fmt.Fprint(file, "hello world\n")

	if err != nil {
		log.Fatal(err)
	}

	readFile, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 12byte格納可能なスライスを用意する
	// 12byteよりオーバーする分は取得できない
	readMessage := make([]byte, 10)

	_, err = readFile.Read(readMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(readMessage))
}

