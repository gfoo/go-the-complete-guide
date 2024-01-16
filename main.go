package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sumup(numbers))

	fmt.Println(sumup2(1, 2, 3, 4, 5))

	fmt.Println(sumup2(numbers...))
}

func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumup2(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
