/*
two ways to check variable's data-type:
- format specifier => %T
- in-built function => reflect.TypeOf
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var message string = "secret message"
	var isCheck bool = true
	var amount float64 = 5466.398

	fmt.Printf("=====Using format specifier %%T=====\n")
	fmt.Printf("%d, %T \n", grades, grades)
	fmt.Printf("%s, %T \n", message, message)
	fmt.Printf("%t, %T \n", isCheck, isCheck)
	fmt.Printf("%.3f, %T \n", amount, amount)

	// using reflect.TypeOf
	fmt.Println("=====using reflect.TypeOf=====")
	fmt.Printf("Type: %v\n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v\n", reflect.TypeOf("John"))
	fmt.Printf("Type: %v\n", reflect.TypeOf(46.123))
	fmt.Printf("Type: %v\n", reflect.TypeOf(false))
}
