package main

import (
	"fmt"
)

func main() {

	const inflationRate = 2.5
	var expenses float64
	var revenue float64
	var taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := calculateEarnBeforeTax(earningsBeforeTax, taxRate)
	ratio := calculateRatio(earningsBeforeTax, earningsAfterTax)

	
	//fmt.Println("Earnings Before Taxation: ", earningsBeforeTax) 
	fmt.Printf("Earnings Before Taxation: %.2f\n", earningsBeforeTax)

	fmt.Println("Earnings After Taxation: ", earningsAfterTax)

	fmt.Println("Profit Ratio: ", ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateEarnBeforeTax(earningsBeforeTax float64, taxRate float64) float64 {
	eat := earningsBeforeTax - (earningsBeforeTax*(taxRate/100))
	return eat
}

func calculateRatio(earningsBeforeTax float64, earningsAfterTax float64) float64 {
	calcRatio := earningsBeforeTax/earningsAfterTax
	return calcRatio
}