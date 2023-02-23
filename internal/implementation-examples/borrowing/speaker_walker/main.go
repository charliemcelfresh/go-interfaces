package main

import (
	"fmt"
)

type SpeakerWalker interface {
	speak()
	walk()
}

type cat struct {
	SpeakerWalker
}

type dog struct{}

func (d dog) speak() {
	fmt.Println("woof!")
}

func (d dog) walk() {
	fmt.Println("walking!")
}

func (c cat) speak() {
	fmt.Println("meow!")
}

func main() {
	d := dog{}
	c := cat{}
	fmt.Println(c.SpeakerWalker) // SpeakerWalker is nil, because it has no underlying concrete object
	//c.walk() panic
	c.speak()
	// cat has no walk method -- it borrows dog's walk method
	c.SpeakerWalker = d
	c.walk()
}
