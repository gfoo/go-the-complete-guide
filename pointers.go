package main

import (
	"fmt"
)

func main() {
	age := 32

	var nilInit *int
	fmt.Println("nilInit:", nilInit)

	// agePointer := &age

	// var agePointer *int
	// agePointer = &age

	var agePointer *int = &age

	fmt.Println("age:", age)
	fmt.Println("age pointer:", agePointer)
	fmt.Println("age pointer val:", *agePointer)

	adultYear := getAdultYears(age)
	fmt.Println("adultYear value:", adultYear)

	adultYearP := getAdultYearsP(&age)
	fmt.Println("adultYear via pointer:", adultYearP)

	adultYearPP := getAdultYearsP(agePointer)
	fmt.Println("adultYear via pointer var:", adultYearPP)

	setAdultYearsP(&age)
	fmt.Println("age directly reset via its pointer:", age)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsP(age *int) int {
	return *age - 18
}

func setAdultYearsP(age *int) {
	*age = *age - 18
}
