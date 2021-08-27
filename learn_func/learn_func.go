package main

import "fmt"

// function declaration
// All function calls in go are called by value i.e when you call a function a value of the variable is passed to the function and not a reference to the variable itself

func doubleFail(a int, arr [2]int, s string {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
}



func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20) //ignore all the values
}

// func main() {
// 	a := addNumbers(2, 3)
// 	fmt.Println(a)

// 	b := addNumbers(3, 2)
// 	fmt.Println(b)

// 	c := addNumbers(4, 2)
// 	fmt.Println(c)
// }

// Go does not have optional parameters or named parameters

func addNumbers(a int, b int) int {
	return a + b
}
