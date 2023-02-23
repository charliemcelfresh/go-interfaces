// Any object that implements sort.Interface will be sorted where passed to sort.Sort()
package main

import (
	"fmt"
	"reflect"
	"sort"
)

type person struct {
	age int
}

type byAge []person

func (b byAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAge) Len() int {
	return len(b)
}

func main() {
	b := byAge{
		{
			age: 11,
		},
		{
			age: 33,
		},
		{
			age: 2,
		},
	}
	fmt.Printf("unsorted   : %v\n", b)
	sort.Sort(b)
	fmt.Printf("sorted asc : %v\n", b)
	r := sort.Reverse(b) // r is borrowing p's Len and Swap methods
	fmt.Printf("type of r  : %v\n", reflect.TypeOf(r))
	sort.Sort(r) // r is using its own Less method, and borrows its embedded object's Swap and Len methods
	fmt.Printf("sorted desc: %v\n", b)
}
