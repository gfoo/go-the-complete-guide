package main

import (
	"fmt"

	"org.gfoo/price-calculator/filemanager"
	"org.gfoo/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		iom := filemanager.News(
			"prices.txt",
			fmt.Sprintf("result_%.0f.json", taxRate*100),
		)
		// iom := cmdmanager.News()
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println(err)
		}
	}
}
