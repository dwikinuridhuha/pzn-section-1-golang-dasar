package main

import "fmt"

func main() {
	// name array
	var names = [...]string{"Alice", "Bob", "Charlie"}

	//slice of names
	var sliceNames = names[1:3] // This will include "Bob" and "Charlie"

	// print the slice
	fmt.Println(names)
	fmt.Println(sliceNames)

	// length, capacity of the slice
	fmt.Println("Length of sliceNames:", len(sliceNames))
	fmt.Println("Capacity of sliceNames:", cap(sliceNames))

	// modifying the slice
	sliceNames[0] = "David"
	fmt.Println("Modified sliceNames:", sliceNames)

	// appending to the slice
	sliceNames = append(sliceNames, "Eve")
	fmt.Println("Slice after appending:", sliceNames)
	fmt.Println("Length of sliceNames after appending:", len(sliceNames))
	fmt.Println("Capacity of sliceNames after appending:", cap(sliceNames))

	// modifying the original array
	names[2] = "Frank"
	fmt.Println("Original names array after modification:", names)

	// modifying the slice again
	sliceNames[1] = "Grace"
	fmt.Println("Modified sliceNames after second modification:", sliceNames)

	// make a new slice with make
	newSlice := make([]string, 2, 5) // length 2, capacity 5
	newSlice[0] = "Frank"
	newSlice[1] = "Grace"
	fmt.Println("New slice created with make:", newSlice)
	fmt.Println("Length of newSlice:", len(newSlice))
	fmt.Println("Capacity of newSlice:", cap(newSlice))

	// appending to the new slice
	newSlice = append(newSlice, "Hank", "Ivy")
	fmt.Println("New slice after appending:", newSlice)
	fmt.Println("Length of newSlice after appending:", len(newSlice))
	fmt.Println("Capacity of newSlice after appending:", cap(newSlice))

	// copying a slice
	copiedSlice := make([]string, len(newSlice), cap(newSlice))
	copy(copiedSlice, newSlice)
	fmt.Println("Copied slice:", copiedSlice)
	fmt.Println("Length of copiedSlice:", len(copiedSlice))
	fmt.Println("Capacity of copiedSlice:", cap(copiedSlice))

}
