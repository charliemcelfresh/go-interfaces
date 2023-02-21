// This is the second half of the code challenge we hand out to prospective Go developers.
// It challenges the developer to implement the pattern that the Go sort package implements:
// create an object that implements an interface, then pass that object into that package's
// operative function.

// Aspiration claims that if someone holds themselves out as a Go developer, but cannot do this problem,
// they have not penetrated one of the basic-intermediate Go patterns, and will not be productive in Go
// here from the outset.

// The sort package requires us to create an object that implements this interface:
// type Interface interface {
//	 Len() int
//	 Less(i, j int) bool
//	 Swap(i, j int)
// }

// Then will sort our object when we pass it to sort.Sort(i Interface)

// The name Interface is intentional, as it describes a particular kind of Go pattern, where the Interface
// describes any object that implements its list (> 1) of methods. As opposed to the -er - named interfaces,
// like Stringer, which implement only one method.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type SkipString struct {
	SkipCounter      int
	InputString      string
	shouldUppercase  map[int]bool
	ValueAsRuneSlice []rune
}

func main() {
	s := NewSkipString(4, "Aspiration.com")
	MapString(&s)
}

func NewSkipString(skipCounter int, s string) SkipString {
	valueAsRuneSlice := make([]rune, len(s))
	shouldUppercase := make(map[int]bool)
	uppercaseCounter := 1
	for i, r := range s {
		if uppercaseCounter%skipCounter == 0 {
			shouldUppercase[i] = true
		}
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			uppercaseCounter++
		}
		valueAsRuneSlice[i] = r
	}
	return SkipString{
		SkipCounter:      skipCounter,
		InputString:      s,
		shouldUppercase:  shouldUppercase,
		ValueAsRuneSlice: valueAsRuneSlice,
	}
}

func (s *SkipString) TransformRune(pos int) {
	if s.shouldUppercase[pos] {
		s.ValueAsRuneSlice[pos] = unicode.ToUpper(s.ValueAsRuneSlice[pos])
	} else {
		s.ValueAsRuneSlice[pos] = unicode.ToLower(s.ValueAsRuneSlice[pos])
	}
}

func (s *SkipString) GetValueAsRuneSlice() []rune {
	var r []rune
	for _, c := range s.InputString {
		r = append(r, c)
	}
	return r
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
	fmt.Println(i)
}

func (s *SkipString) String() string {
	sl := []string{}
	for _, r := range s.ValueAsRuneSlice {
		sl = append(sl, string(r))
	}
	return strings.Join(sl, "")
}
