package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var investAmount = 1000
	// var expectReturnRate = 5.5
	// var years = 10
	// var futureValue = float64(investAmount) * math.Pow(1+expectReturnRate/100, float64(years))

	// var investAmount float64 = 1000
	// var expectReturnRate = 5.5
	// var years float64 = 10
	// var futureValue = investAmount * math.Pow(1+expectReturnRate/100, years)

	// const inflationRate = 6.5

	expectReturnRate, years := 5.5, 10.0
	var investAmount float64

	outputText("Enter invest amount")
	fmt.Scan(&investAmount)

	// futureValue := investAmount * math.Pow(1+expectReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calcFutureValues(investAmount, expectReturnRate, years)

	// fmt.Println("Future value:", futureValue)
	// fmt.Printf("Future real value: %6.2f\n", futureRealValue)

	fmt.Printf(`
	Future value: %f
	Future real value: %6.2f
	`, futureValue, futureRealValue)

}

func outputText(text string) {
	fmt.Println(text)
}

//	func calcFutureValues(investAmount, expectReturnRate, years float64) (float64, float64) {
//		fv := investAmount * math.Pow(1+expectReturnRate/100, years)
//		frv := fv / math.Pow(1+inflationRate/100, years)
//		return fv, frv
//	}
func calcFutureValues(investAmount, expectReturnRate, years float64) (fv float64, frv float64) {
	fv = investAmount * math.Pow(1+expectReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}
