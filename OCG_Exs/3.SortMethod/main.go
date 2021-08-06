package main

import (
	"fmt"
)

func main() {
	array := [6]int{2, 3, 4, 5, 6, 7}
	findMax(array[:])
	findMin(array[:])
	calculateAverage(array[:])
}

func findMax(arr []int) {

	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

func findMin(arr []int) {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}

func calculateAverage(arr []int) {
	ave := float64(0)
	total := 0
	for _, v := range arr {
		total = total + (v)
		ave = float64(total) / float64(len(arr)) // ave is needs to be of type float
	}
	fmt.Println(ave)
}
