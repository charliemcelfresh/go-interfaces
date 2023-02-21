// This is a simplified version of the fmt package.
// If an object has a String() method, it implements the fmt package's Stringer interface.
// If an object does not have a String() method, the fmt package tries very hard to print
// the string version of that object.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// StringerType has a String() method
type StringerType struct {
	Name string
}

func (s StringerType) String() string {
	return s.Name
}

// NotStringerType has no String() method
type NotStringerType struct {
	Name string
}

type Stringer interface {
	String() string
}

func main() {
	// StringerType
	s := StringerType{Name: "StringerType"}
	fmt.Println(ToString(s))

	// int
	fmt.Println(ToString(1))

	// NonStringerType
	n := NotStringerType{Name: "NotStringerType"}
	fmt.Println(ToString(n))
}

func ToString(any interface{}) string {
	// switch statements can switch on both interface and concrete types
	switch v := any.(type) { // grab the concrete type from any
	case Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	default:
		return reflect.ValueOf(any).String()
	}
}
