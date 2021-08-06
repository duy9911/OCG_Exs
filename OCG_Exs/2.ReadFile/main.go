package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	readFile("duy.txt")
}

func readFile(filename string) {
	arr := []int{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range strings.Fields(string(data)) { // - strings.Fields convert string data to string []array
		num, _ := strconv.Atoi(v) // -convert string to int
		arr = append(arr, num)    //-import values to []arr
	}
	findMax(arr)
	findMin(arr)
	calculateAverage(arr)
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
