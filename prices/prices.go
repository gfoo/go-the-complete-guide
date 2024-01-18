package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open the file prices.csv")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file prices.csv")
		fmt.Println(err)
		file.Close()
		return
	}
	prices := make([]float64, len(lines))
	for i, line := range lines {
		prices[i], err = strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing the file prices.csv")
			fmt.Println(err)
			file.Close()
			return
		}
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string, len(job.InputPrices))
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", (price * (1 + job.TaxRate)))
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
