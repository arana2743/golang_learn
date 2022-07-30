/*
Type casting :
1. int to float - can be easily done
2. float to int - float value gets truncated to convert to int
3. bool to int/float - cannot happen

strconv() package
1. Itoa() - convert integer to string
2. Atoi() - converts string to integer,
	returns 2 values : the corresponding integer, error (if any)
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ifv int = 259
	var f_ifv float64 = float64(ifv)
	fmt.Printf("integer: %d, converted float: %.2f \n", ifv, f_ifv)

	var fi float64 = 44.95
	var i_fi int = int(fi)
	fmt.Printf("float: %.2f, converted to int: %d \n", fi, i_fi)

	// strconv()
	fmt.Println("====strconv package functions====")
	// Itoa()
	var num1 int = 199
	var str_num1 string = strconv.Itoa(num1)
	fmt.Println("strconv.Itoa() - int to string")
	fmt.Printf("%q \n", str_num1)

	// Atoi()
	num2, err := strconv.Atoi("yes")
	fmt.Println("strconv.Atoi() - string to int")
	fmt.Printf("%v %v ", num2, err)
}
