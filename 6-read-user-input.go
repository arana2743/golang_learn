/*
Scanf - used to read input
Scanf returns two values
 - count => count of input read
 - err => error while reading any input
*/
package main

import "fmt"

func main() {
	var a string
	var b int
	fmt.Print("Enter a string and a number(space separated): ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Println(a, b)
}
