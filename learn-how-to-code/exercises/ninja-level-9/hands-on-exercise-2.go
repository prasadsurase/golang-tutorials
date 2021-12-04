package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func (p *person) speak() {
	fmt.Println("I am " + p.firstname + " " + p.lastname)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{firstname: "Prasad", lastname: "Surase"}
	saySomething(&p)
}
