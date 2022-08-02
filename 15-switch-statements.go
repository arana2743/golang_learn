/*
Switch statements in golang have additional concept:
1. Fallthrough
	- used to force the execution flow to fall through the successive case block
	- meaning very next case block also gets executed (even if it's not a match) after the matched case.
2. Switch statements can also contain conditions

*/
package main

import "fmt"

func main() {
	fmt.Println("====Switch statements====")
	var i int = 500
	switch i {
	case 10:
		fmt.Println("i is 10.")
		fallthrough // next case block will also execute when this is matched
	case 100, 200:
		fmt.Println("i is either 100 or 200")
	case 500:
		fmt.Println("i is 500")
		fallthrough // next case block will also execute when this is matched
	default:
		fmt.Println("i is something else (other than 10, 100, 200)!")
	}

	var a, b int = 10, 20
	switch {
	case a+b == 30:
		fmt.Println("equal to 30")
	case a+b <= 30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}
}
