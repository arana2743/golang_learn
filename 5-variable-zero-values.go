/*
Variables when declared but no initialised have default value set in Golang
int - 0
float - 0.0
bool - false
string - ""
pointers, functions, interfaces, maps = nil
*/
package main

import "fmt"

func main() {
	var (
		zeroInt    int
		zeroFloat  float64
		zeroBool   bool
		zeroString string
	)
	fmt.Println(zeroInt, zeroFloat, zeroBool, zeroString, ".")

}
