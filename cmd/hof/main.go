package main

import (
	"fmt"
)

// this is a "Higher Order function, it's a function, that takes argument as a function"
func filter(numbers []int, f func(int) bool) []int {
	// highlight: we declare a new array here, which allows us to create pure functions
	// immutability -> numbers array accepted as an argument, wont be changed, we create a new one and return that
	var result []int 
	for _ ,value := range(numbers) {
		if f(value) {
			result = append(result, value)
		}
	}
	return  result
}

func isEven(num int) bool {
	return num % 2 == 0
} 

func main(){
	values := []int{1,2,3,4,5,6}
	result := filter(values, isEven)
	fmt.Println(result)
}