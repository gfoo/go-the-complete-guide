package prices

import (
	"fmt"

	"org.gfoo/price-calculator/conversion"
	"org.gfoo/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string, len(job.InputPrices))
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", (price * (1 + job.TaxRate)))
	}
	job.TaxIncludedPrices = result
	filemanager.WriteJSON(job, fmt.Sprintf("result_%.0f.json", job.TaxRate*100))
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
