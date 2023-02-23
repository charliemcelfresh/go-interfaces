// direct assignment
package main

import (
	"fmt"
)

type cat struct{}

// mover contains the cat method move(), but not speak()
type mover interface {
	move()
}

func main() {
	// Create a cat, and call its methods
	c := cat{}

	c.move()
	c.speak()
	// assign the concrete type cat to the mover interface
	var m mover = c

	m.move()
	//m.speak() // m is a mover, not a speaker

	assertCat := m.(cat) // type assertion provides access to the underlying concrete type
	assertCat.move()
	assertCat.speak()
}

func (c cat) speak() {
	fmt.Println("meow")
}

func (c cat) move() {
	fmt.Println("slink")
}
