/*
Slices
- slices are more flexible than arrays since they are variable types (elements can be added and removed)
- during slice creation => compiler first creates a underlying array and then returns a slice

- also when slice is created from underlying existing array then, any change in elements in slice gets reflected to underlying array as well
Since slice is nothing but a reference to the existing array
*/
package main

import (
	"fmt"
)

func main() {
	slice_grades := []int{10, 20, 30, 40}
	fmt.Println(slice_grades)
	fmt.Printf("len=%v, cap=%v, type=%T \n", len(slice_grades), cap(slice_grades), slice_grades)

	slice_grades = append(slice_grades, 50, 60, 10)
	fmt.Println(slice_grades)

	// append another slice to slice
	num_arr := [...]int{10, 20}
	num_slice := num_arr[:]
	fmt.Printf("Before: len=%v cap=%v \n", len(slice_grades), cap(slice_grades))
	slice_grades = append(slice_grades, num_slice...)
	fmt.Printf("After: len=%v cap=%v \n", len(slice_grades), cap(slice_grades))

	// copying a slice to another
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice) // returned the number of elements copied
	fmt.Println(num)
	fmt.Println(dest_slice)

	// creating a slice from already existing array
	suspects_arr := [...]string{"Kang", "Eternals", "Celestials", "Thanos", "Eros", "Loki", "Skrulls"}
	suspects_slice := suspects_arr[2:4]
	fmt.Println(suspects_arr)
	fmt.Println(suspects_slice)

	// another way of declaring a slice
	slice_new := make([]int, 10, 12)
	fmt.Println(slice_new)
}
