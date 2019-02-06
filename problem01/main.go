package main

import "fmt"

func main() {

	sum := sumOfMultiplesOf(1000, []int{3, 5})
	fmt.Printf("Sum is %d\n", sum)
}

func sumOfMultiplesOf(limit int, numbers []int) int {
	//
	result := 0

	for i:=1; i < limit; i++ {
		if isMultiple(i, numbers) {
			result += i
		}
	}

	return result
}

func isMultiple (number int, numbers []int) bool {

	for _, n := range numbers {
		if number % n == 0 {
			return true
		}
	}

	return false
}
