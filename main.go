package main

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)
	doubled := tranformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := tranformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func tranformNumbers(numbers *[]int, transform transformFn) []int {
	doubleNumbers := []int{}
	for _, v := range *numbers {
		doubleNumbers = append(doubleNumbers, transform(v))
	}
	return doubleNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
