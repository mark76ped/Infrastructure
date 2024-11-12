package profitReimagined

import (
	"fmt"
	"os"
)


const calculateFinancialsFile = "calculated_financials.txt"


func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func writeEbtToFile(calculateEbt, calculatedProfit, calculatedRatio float64) {
	calculateFinancialsText := fmt.Sprintf("EBT %.1f\nProfit: %.1f\nRatio: %.3f\n", calculateEbt, calculatedProfit, calculatedRatio)
	os.WriteFile(calculateFinancialsFile, []byte(calculateFinancialsText), 0644)
}


func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	if revenue <= 0 {
		panic("Revenue cannot be nothing.")
	} else if expenses <= 0 {
		panic("Expenses cannot be nothing.")
	} else if taxRate <= 0 {
		panic("The Gov't would never make taxes 0 or less for you. You're just a small business.")
	}

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	writeEbtToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
