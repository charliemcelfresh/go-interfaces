// direct assignment
package main

import "fmt"

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
	var i mover = c

	i.move()
	// i.speak() // speak() method is not available to i

	castCat := i.(cat) // type assertion provides access to the underlying concrete type
	castCat.move()
	castCat.speak()

	if castCat, ok := i.(cat); ok { // perform checks using the "ok" idiom
		castCat.move()
		castCat.speak()
	}
}

func (c cat) speak() {
	fmt.Println("meow")
}

func (c cat) move() {
	fmt.Println("slink")
}
