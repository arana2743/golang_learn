/*
Format specifiers:
%v - formats the value in default format
%d - formats the decimal integers
%c - character
%s - plain string
%t - boolean (true or false)
%f - floating numbers
%.2f - floating number upto 2 decimal places
%q - to format character/string into quoted characters/string
%T - data_type of the value
*/
package main

import "fmt"

func main() {
	var name string = "John Constantine"
	fmt.Println("Hello there,", name) // remember "," already add a space

	// now with Printf
	var location string = "New Delhi"
	fmt.Printf("Nice to see you here, at %v", location)
}
