package main

import "fmt"
import "github.com/joelfransson/packagetest/p1"

type person struct {
	name string
}

func (p person) SayHello() {
	s := fmt.Sprintf("Hello %s", p.name)
	fmt.Println(s)
}

func main() {
	p := person{name: "kalle"}
	p1.SayHello(p)
}
