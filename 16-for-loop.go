/*
Golang has only for-loop,
however for-loop can also be written in form of while loop as well.
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	// use of break statement in for-loop
	for i := 1; i <= 10; i++ {
		if i >= 5 {
			break
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// use of continue statement in for-loop
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%v ", i)
	}
}
