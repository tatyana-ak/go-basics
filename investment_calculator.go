package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000.0
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formatterFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formatterRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	// fmt.Print(formatterFV, formatterRFV)

	fmt.Printf(`Future value: %.2f
	Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}
