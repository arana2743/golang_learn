/*
Block scopes
- inner block can access variables declared withing the outer scope
- however outer blocks cannot access variables withing inner blocks
** that's what usually happens in every programming language :lol
*/
package main

import "fmt"

// global variable : declared outside of any block/function
var globalVar string = "Global Variable"

func main() {
	city := "London"
	{
		country := "England"
		city := "Kolkata"
		fmt.Println(city)
		fmt.Println(country)
	}
	// fmt.Println(country) // this will fail
	fmt.Println(city)
}
