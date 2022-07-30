package main

import "fmt"

func main() {
	var name string
	name = "John"
	fmt.Printf("The name is %s Cena.", name)

	// declaring multiple variables of same type
	var s, t string
	s = "foo"
	t = "doo"
	fmt.Printf("\ns = %s, t = %s", s, t)

	// declaring multiple variable of different types
	var (
		d   string
		age int
	)
	d = "Done"
	age = 20
	fmt.Printf("\nd = %s, age = %d", d, age)

	// shorthand way of variable declaration and initialisation
	short := "Shorthand notation"
	fmt.Printf("\nshort = %s and it's type is: %T", short, short)
}
