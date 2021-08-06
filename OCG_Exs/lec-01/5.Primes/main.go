package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	checkPrimeFile("duy.txt")
}

func checkPrimeFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range strings.Fields(string(data)) {
		num, _ := strconv.Atoi(v)
		if IsPrime(num) == true {
			fmt.Println("Is prime number:", num)
		} else {
			fmt.Println("Not prime number:", num)
		}
	}
}

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
