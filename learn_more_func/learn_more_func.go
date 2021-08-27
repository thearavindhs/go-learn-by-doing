package main

import "fmt"

// func addOne(a int) int {
// 	return a + 1
// }

// func main() {
// 	myAddOne := addOne //referecing another function i.e assigning to a variable
// 	fmt.Println(addOne(1))
// 	fmt.Println(myAddOne(1))
// }

// func main() {

// 	// This will cause compilation error, you need to create an ananoymous function and assign to the variable
// 	// func addOne(a int) int {
// 	// 	return a + 1
// 	// }

// 	myAddOne := func(a int) int {
// 		return a + 1
// 	}
// 	fmt.Println(myAddOne(1))
// }

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func main() {
	printOperation(1, addOne)
	printOperation(1, addTwo)
}
