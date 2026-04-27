package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000.0
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formatterFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formatterRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	// fmt.Print(formatterFV, formatterRFV)

	fmt.Printf(`Future value: %.2f
	Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)
}
