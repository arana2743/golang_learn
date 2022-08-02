/*
Arrays:
- fixed length (size once set cannot be changed)
- elements in array should be of same data-type

array properties:
- length
- capacity (in case of arrays length and capacity is same)
- contains pointer to first element
*/
package main

import "fmt"

func main() {
	var grades [5]int = [5]int{10, 20, 30, 40, 100}
	var fruits [3]string
	prices := [3]float64{10.22, 22.35, 100.98}
	suspects := [...]string{"Jonh D", "Joker", "Riddler"}

	fmt.Printf("length=%v, capacity=%v \n", len(grades), cap(grades))
	fmt.Println(grades)
	fmt.Println(fruits)
	fmt.Println(prices)
	fmt.Println(suspects)

	// looping through arrays
	// way1
	for i := 0; i < len(suspects); i++ {
		fmt.Printf("%v ", suspects[i])
	}
	fmt.Println()
	// way2
	for index, element := range grades {
		fmt.Printf("element[%v]: %v \n", index, element)
	}
}
