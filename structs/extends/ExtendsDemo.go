package main

import "fmt"

type Person struct {
	name string
	id   int
}

func (person *Person) sayHello() {
	fmt.Printf("hello, i'm %s\n", person.name)
	person.name = "changedName"
}

func (person *Person) introduce() {
	fmt.Printf("my name is %s\n", person.name)
}

type Female struct {
	Person
}

func (female *Female) sayHello() {
	female.Person.sayHello()
	fmt.Println("im done")
}

func main() {
	female := Female{Person{"name", 12}}
	female.sayHello()
	female.introduce()
}
