package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{1.2, 3.4}

	fmt.Println(prices)

	updatedPrices := append(prices, 5.6)
	fmt.Println(updatedPrices)

	prices[0] = 0.0
	fmt.Println(prices, updatedPrices)

	discountPrices := []float64{0.5, 0.6, 0.7}
	prices = append(prices, discountPrices...)
	fmt.Println(prices, discountPrices)

}

// func main() {

// 	prices := [5]float64{1.2, 3.4, 5.6, 7.8, 9.0}

// 	fmt.Println(prices)

// 	var productNames [5]string = [5]string{"Pen", "Pencil", "Marker"}
// 	productNames[3] = "Eraser"

// 	fmt.Println(productNames)
// 	fmt.Println(productNames[0])
// 	fmt.Println(productNames[1:3])
// 	fmt.Println(productNames[:2])
// 	fmt.Println(productNames[2:])

// 	// pointer
// 	featuredProduct := productNames[:2]
// 	fmt.Println(featuredProduct)
// 	featuredProduct[0] = "NewPen"
// 	fmt.Println(productNames)
// 	fmt.Println(featuredProduct)

// 	fmt.Println(len(productNames), cap(productNames))
// 	fmt.Println(len(featuredProduct), cap(featuredProduct))

// 	featuredProduct = featuredProduct[:cap(featuredProduct)]
// 	fmt.Println(featuredProduct)
// 	fmt.Println(len(featuredProduct), cap(featuredProduct))
// }
