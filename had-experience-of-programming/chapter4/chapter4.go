package chapter4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	doJsonFiles()

	doIoUtil()
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
	defer readFile.Close()

	// ここが重要
	// 10byte格納可能なスライスを用意する
	// []byte を必要な長さを用意してあげないといけない
	// 10byteよりオーバーする分は取得できない
	readMessage := make([]byte, 10)

	_, err = readFile.Read(readMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(readMessage))
}

func doJsonFiles() {
	person := &Person{
		Id:      1,
		Name:    "Goper",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	// ここからは
	// JSONの書き込み
	// ファイルを開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// エンコーダーの取得
	encoder := json.NewEncoder(file)

	// ファイルへの書き出し(Encodeを使う)
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}

	// ここからは
	// JSONの読み込み
	readFile, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	// データを読み込む変数
	var readPerson Person

	// デコーダの取得
	decoder := json.NewDecoder(readFile)
	// JSONデコードしたデータの書き込み(Decodeを使う)
	err = decoder.Decode(&readPerson)
	if err != nil {
		log.Fatal(err)
	}

	// 読みだした結果の表示
	fmt.Println(readPerson)
}

func doIoUtil() {
	// ReadAll
	file, _ := os.Open("./file.txt")
	message, _ := ioutil.ReadAll(file)
	fmt.Println(string(message))
	// Printfだと[]byteもそのまま出力できる
	//fmt.Printf("%s\n", message)

	// WriteFile
	message = []byte("hello world???")
	// permを0777としているが実際作成されるのは755のファイルumaskが効いているのかな
	ioutil.WriteFile("./file_io_util.txt", message, 0777)

	// ReadFile
	message, _ = ioutil.ReadFile("./file_io_util.txt")
	fmt.Println(string(message))

}
