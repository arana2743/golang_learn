/*
Constants - can be declared with or without declaring data-type
hence gets defined as 2 types:
 1. untyped => wherein we don't declare the type of constant variable
 2. types => whereing we declare the type of constant variable

Also constants need to be declared and initialised at the same time
And shorthand declartion (:=) is also not allowed for constants
*/

package main

import "fmt"

func main() {
	const name string = "John"
	fmt.Printf("%v, %T", name, name)
}
