/*
maps are dictionaries (key-value pairs)
*/
package main

import "fmt"

func main() {
	// declaring a nil map - in this we cannot add new values
	var my_map map[string]int
	fmt.Println(my_map)

	// declaring and initialising a map
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)

	// using make function
	my_map2 := make(map[string]int)
	fmt.Println(my_map2)

	// add new key-value pairs
	my_map2["one"] = 1
	my_map2["two"] = 2
	my_map2["hundred"] = 100
	my_map2["three"] = 3

	fmt.Println(my_map2)

	// deleting a key-value pair
	delete(my_map2, "hundred")

	fmt.Println(my_map2)
	fmt.Printf("lenght=%v, type=%T \n", len(my_map2), my_map2)

	// getting value for a key
	value, found := my_map2["three"]
	fmt.Printf("key found = %v, key value = %v\n", found, value)

	for key, value := range my_map2 {
		fmt.Printf("map[%v] = %v \n", key, value)
	}

}
