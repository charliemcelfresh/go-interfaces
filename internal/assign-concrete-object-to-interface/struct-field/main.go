// struct field is type mover, not type cat
package main

import "fmt"

type petStore struct {
	cat mover
}

type cat struct{}

// mover contains the cat method move(), but not speak()
type mover interface {
	move()
}

func main() {
	c := cat{}
	z := petStore{
		cat: c,
	}
	z.describe()
}

func (p petStore) describe() {
	p.cat.move()
	//p.cat.speak() // cat is a mover, and speak is not in the mover interface
}

func (c cat) speak() {
	fmt.Println("meow")
}

func (c cat) move() {
	fmt.Println("slink")
}
