package main

import (
	"fmt"

	"org.gfoo/price-calculator/filemanager"
	"org.gfoo/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		iom := filemanager.News(
			"prices.txt",
			fmt.Sprintf("result_%.0f.json", taxRate*100),
		)
		// iom := cmdmanager.News()
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate)

		go priceJob.Process(doneChans[index], errorChans[index])
		// err := priceJob.Process()
		// if err != nil {
		// 	fmt.Println(err)
		// }

	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}

}
