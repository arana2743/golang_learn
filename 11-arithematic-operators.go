/*
Arithematic operators are:
"+" - addition
"-" - subtraction
"/" - division
"*" - multiplication
"%" - modulus
"++", "--" - increment and decrement (and both are only post increments/decrements only )
*/
package main

import "fmt"

func main() {
	fmt.Println(10 + 20 - 10)
	fmt.Println(20/10, 32/10, 32%10)
	fmt.Println(20 * 44)
	value := 10
	value++
	fmt.Println(value)

	var x, y int = 100, 9
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)
}
