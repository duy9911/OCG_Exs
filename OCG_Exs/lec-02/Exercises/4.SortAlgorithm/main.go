package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort ne`")
	inputArray := []int{1, 2, 3, 4, 5, 5, 67, 8, 8}
	outputArray := bubbleSort(inputArray)
	fmt.Println(outputArray)
}

func bubbleSort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			c := input[i+1]
			input[i+1] = input[i]
			input[i] = c
		}
	}
	return input
}
