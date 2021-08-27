package main

import "fmt"

// func main() {
// 	a := 10
// 	if b := 10; a > 5 {
// 		fmt.Println("a is bigger than 5", b)
// 	} else {
// 		fmt.Println("a is lesser than 5", b)
// 	}
// 	fmt.Println(a)
// }

func main() {
	// for i := 0; i < 10; i++ { // for  statement can be turned into a while statement without the condition and increment
	// for k, v := range s {
	//    fmt.Println(k,v, string(v))
	//}
	i := 3
	for i > 0 {
		if i == 5 {
			break
		} else {
			fmt.Println(i)
			i += 1
		}
	}
}
