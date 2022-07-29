/*
Data types in Golang:
	1. String (16 bytes)
	2. Numbers:
		Integers (signed and unsigned from 1-8 bytes each type)
		Float (4 and 8 bytes)
	3. Boolean (1 byte)
	4. Arrays
	5. Slices
	6. Maps
*/
package main

import "fmt"

func main() {
	title := "Doctor Sleep"
	fmt.Println(title)

	var value int = 100
	var price float64 = 77.832
	fmt.Println(value)
	fmt.Println(price)

	var name string = "John D"
	fmt.Println(name)

	var isValid bool = false
	fmt.Println(isValid)
}
