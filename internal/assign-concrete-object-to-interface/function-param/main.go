// struct field is type mover, not type cat
package main

import "fmt"

type cat struct{}

// movers cannot speak
type mover interface {
	move()
}

func main() {
	c := cat{}
	describe(c)
}

func describe(cat mover) {
	cat.move()
	// cat.speak() // cat is a mover, and cannot speak
}

func (c cat) speak() {
	fmt.Println("meow")
}

func (c cat) move() {
	fmt.Println("slink")
}
