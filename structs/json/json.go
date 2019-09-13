package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func struct2json() {
	var user = User{"wzy", 22}
	target, _ := json.Marshal(user)
	fmt.Println(string(target))
}

func json2struct() {
	var jsonStr = `
        {
                "name":"wzy",
                "age":12
        }
        `
	var user User
	_ = json.Unmarshal([]byte(jsonStr), &user)
	fmt.Println(user)
}

func main() {
	struct2json()
	json2struct()
	map2json()
	json2map()
}

func map2json() {
	obj := map[string]interface{}{"name": "wzy", "age": 12}
	target, _ := json.Marshal(obj)
	fmt.Println(string(target))
}

func json2map() {
	var jsonStr = `
        {
                "name":"wzy",
                "age":12
        }
        `
	obj := map[string]interface{}{}
	_ = json.Unmarshal([]byte(jsonStr), &obj)
	fmt.Println(obj)
}
