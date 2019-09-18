package main

import "fmt"

type User struct {
	Pwd  string
	Name string
	UID  int64
}

func (user User) SayHello() {
	fmt.Println("hello my name is " + user.Name)
}

func test(user *User) {
	user.Name = "newName"
}

func main() {
	var u = User{Name: "q2", Pwd: "wangzy"}
	fmt.Println(u)
	test(&u)
	fmt.Println(u)
}
