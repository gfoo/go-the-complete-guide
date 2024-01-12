package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	// usrName := []string{}
	//
	usrName := make([]string, 2, 3)

	usrName[0] = "John1"
	fmt.Println(usrName)

	usrName = append(usrName, "John2")
	usrName = append(usrName, "Doe")

	fmt.Println(usrName)

	// courseRating := map[string]float64{}
	// courseRating := make(map[string]float64, 2)
	courseRating := make(floatMap, 2)

	courseRating["GoLang"] = 5
	courseRating["Python"] = 4.5
	courseRating["Java"] = 4.5

	// fmt.Println(courseRating)
	courseRating.output()

	for index, value := range usrName {
		fmt.Println(index, value)
	}

	for key, value := range courseRating {
		fmt.Println(key, value)
	}

}
