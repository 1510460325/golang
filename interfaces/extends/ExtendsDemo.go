package main

import "fmt"

type Person struct {
	name string
	id   int
}

func (person Person) sayHello() {
	fmt.Printf("hello, i'm %s\n", person.name)
}

type Famale struct {
	Person
}

func (famale Famale) sayHello() {
	famale.Person.sayHello()
	fmt.Printf("repeat it")
}

func main() {
	person := Person{"abc", 1}
	person.sayHello()
	female := Famale{Person{"name", 12}}
	female.sayHello()
}
