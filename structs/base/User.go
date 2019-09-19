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
	change()
}

func change() {
	data := []int{1, 3, 2, 4, 2, 1}
	ans := -1
	for i := 0; i < len(data); i++ {
		for data[i] != i {
			if data[data[i]] == data[i] {
				ans = data[i]
				break
			}
			j := data[i]
			a := data[i]
			data[i] = data[j]
			data[j] = a
		}
	}
	fmt.Println(ans)
}
