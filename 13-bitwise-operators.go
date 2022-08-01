/*
Bitwise operators are:
"&" - Bitwise AND
"|" - Bitwise OR
"^" - Bitwise XOR (the result of XOR is 1 if 2 bits are opposite)
">>" - Right shift
"<<" - Left shift
*/
package main

import "fmt"

func main() {
	fmt.Println("====Bitwise operators====")
	var x, y int = 12, 25
	b_and := x & y // bitwise and
	b_or := x | y  // bitwise or
	b_xor := x ^ y // bitwise xor
	fmt.Println(b_and, b_or, b_xor)

	var num int = 212
	left_shift := num << 1  // left shift
	fmt.Println(left_shift) // left shift by 1 doubles the decimal number

	right_shift := num >> 1  // right shift
	fmt.Println(right_shift) // right shift by 1 half's the decimal number

}
