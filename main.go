package main

import (
	"fmt"
)

func filter(numbers []int, f func(int) bool) []int {
	var result []int
	for _, value := range numbers {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func isEven(value int) bool {
	return value % 2  == 0
}


func main() {
	myNumbers := []int{1,2,3,4,5,6}
	newNumbers := filter(myNumbers, isEven)
	fmt.Println(newNumbers)
}
