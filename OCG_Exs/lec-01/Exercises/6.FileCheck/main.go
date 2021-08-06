package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("duy.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Contents of file:", string(data))
	checkValid(string(data))
}

func checkValid(dataFile string) {
	dataDummy := 4

	for _, v := range strings.Fields(string(dataFile)) {
		number, _ := strconv.Atoi(v)
		count := 0
		if number == dataDummy {
			fmt.Printf("\n%v equal %v", number, dataDummy)
			count++
		} else {
			fmt.Println("No value equal")
		}
		fmt.Println("\nTotal: ", count)
	}
}
