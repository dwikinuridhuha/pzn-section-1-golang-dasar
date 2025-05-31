package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	var numbers = [5]int{1, 2, 3, 4, 5}
	var matrix = [...]string{
		"one", "two", "three",
	}

	fmt.Println("Names:", names)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Matrix:", matrix)
}
