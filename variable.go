package main

import "fmt"

func main() {

	var singleName string

	var (
		names   [3]string
		numbers = [5]int{1, 2, 3, 4, 5}
		matrix  = [...]string{
			"one", "two", "three",
			"four", "five",
			"six", "seven", "eight",
			"nine", "ten",
		}
	)

	lengthNames := len(names)
	lengthNumbers := len(numbers)
	lengthMatrix := len(matrix)

	// Assign values to the singleName & names array
	singleName = "Alice"
	names[0] = "Tom"
	names[1] = "Jerry"
	names[2] = "Spike"

	// Output the arrays and their lengths
	fmt.Println("Single Name:", singleName)
	fmt.Println("Single Name index 3:", singleName[3])
	fmt.Println("Names:", names)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Matrix:", matrix)
	fmt.Println("Length of names array:", lengthNames)
	fmt.Println("Length of numbers array:", lengthNumbers)
	fmt.Println("Length of matrix array:", lengthMatrix)
}
