package main

import "fmt"

func main() {
	// map example
	var names = map[string]string{
		"first":  "Alice",
		"second": "Bob",
		"third":  "Charlie",
	}

	// Accessing values in the map
	fmt.Println("First name:", names["first"])
	fmt.Println("Second name:", names["second"])
	fmt.Println("Third name:", names["third"])

	// Adding a new key-value pair
	names["fourth"] = "David"
	fmt.Println("After adding fourth name:", names)

	// Modifying an existing key-value pair
	names["second"] = "Eve"
	fmt.Println("After modifying second name:", names)

	// Deleting a key-value pair
	delete(names, "third")
	fmt.Println("After deleting third name:", names)
}
