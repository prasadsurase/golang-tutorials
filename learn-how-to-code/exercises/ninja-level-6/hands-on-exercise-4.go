package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func (p *person) speak() {
	fmt.Println("I am", p.firstname, p.lastname, "and I am", p.age, "years old.")
}

func main() {
	p1 := person{
		firstname: "James", lastname: "Bond", age: 42,
	}

	p1.speak()
}
