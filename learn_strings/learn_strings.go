package main

import (
	"fmt"
)

func main() {
	// emtpy string of length 0, if value is not assigned its defaulted to 0
	// var s string
	// s = "Hello World!"
	s := `"Hello, World!"` // use string literal
	fmt.Println(s)
	b := s[2:6]
	fmt.Println(b)
	a := s[2]
	fmt.Println(string(a)) // Go treats string as unmodifiable sequence of bytes
	const y = "id"
	fmt.Println(y)
	// y = "hello" // cant reassign constant

}
