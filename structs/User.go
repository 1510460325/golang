package main

import "fmt"

type User struct {
	Pwd  string
	Name interface{}
}

func main() {
	var u = User{Name: "q2", Pwd: "wangzy"}
	fmt.Println(u)
}
