package main

import (
	"fmt"

	"example.com/calculator/cmdmanager"
	"example.com/calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*10))
		cmds := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmds, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not proccess job", err)

		}
	}

}
