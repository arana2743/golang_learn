package main

import "fmt"

func main() {
	fmt.Println("====If-elseif-else statements====")
	var a string = "Sad"
	if a == "Happy" {
		fmt.Printf("I am %v, I will go for a run.\n", a)
	} else if a == "Sad" {
		fmt.Printf("I am %v, I will watch comedy today to cheer myself.\n", a)
	} else {
		fmt.Printf("Oh no! I don't know how I feel today!\n")
	}
}
