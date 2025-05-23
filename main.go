package main

import (
	"github.com/AntonioGuilhermeDev/pricecalculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludePriceJob(taxRate)
		priceJob.Process()
	}
}
