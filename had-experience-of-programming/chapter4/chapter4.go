package chapter4

import (
	"encoding/json"
	"fmt"
	"log"
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
