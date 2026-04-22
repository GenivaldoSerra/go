package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)

	// Append an element to the array
	arr = append(arr, 6)
	fmt.Println("After appending 6:", arr)

	// Remove the last element from the array
	arr = arr[:len(arr)-1]
	fmt.Println("After removing the last element:", arr)

	// Accessing elements in the array
	fmt.Println("First element:", arr[0])
	fmt.Println("Second element:", arr[1])

	// Iterating over the array
	fmt.Println("Iterating over the array:")
	for i, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}
}