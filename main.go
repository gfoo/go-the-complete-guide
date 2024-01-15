package main

import "fmt"

type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	moreNumbers := []int{5, 1, 2}

	transformFn1 := getTransformFn(&numbers)
	transformFn2 := getTransformFn(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func getTransformFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
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
