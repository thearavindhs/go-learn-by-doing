package main

import (
	"fmt"
	"os"
)

// func main() {
// 	word := os.Args[1]

// 	if word == "hello" {
// 		fmt.Println("Hi Yourself")
// 	} else if word == "hi" {
// 		fmt.Println("Hello Yourself")
// 	} else if word == "see" {
// 		fmt.Println("bye")
// 	} else {
// 		fmt.Println("gone")
// 	}
// }

func main() {
	word := os.Args[1]
	greet := "greetings"

	switch l := len(word); word {
	case "hi":
		fmt.Println("Very Informal")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell":
	case "goodbye", "bye":
		fmt.Println("So long")
	case greet:
		fmt.Println("Salutations")
	default:
		fmt.Println("I dont know what you said", l)
	}
}
