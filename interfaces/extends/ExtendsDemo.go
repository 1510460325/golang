package main

import "fmt"

type Person struct {
	name string
	id   int
}

func (person Person) sayHello() {
	fmt.Printf("hello, i'm %s\n", person.name)
}

type Female struct {
	Person
}

func (female Female) sayHello() {
	female.Person.sayHello()
	fmt.Printf("repeat it")
}

func main() {
	person := Person{"abc", 1}
	person.sayHello()
	female := Female{Person{"name", 12}}
	female.sayHello()
}
