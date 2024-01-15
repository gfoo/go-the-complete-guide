package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(val int) int {
		return val * 2
	})
	fmt.Println(transformed)

	double := createTrans(2)
	triple := createTrans(3)

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTrans(factor int) func(int) int {
	return func(val int) int {
		return val * factor
	}
}
