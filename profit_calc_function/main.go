package main

import "fmt"

func main() {
	revenue := takeInputs("Revenue: ")
	expences := takeInputs("Expences: ")
	taxRate := takeInputs("tax Rate: ")

	PBT, PAT := profitBeforeAndAfterTax(revenue, expences, taxRate)

	fmt.Printf("Profit before tax: %.1f\n", PBT)
	fmt.Printf("Profit After Tax: %.1f", PAT)
}

func takeInputs(inputText string) float64 {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)
	return userInput
}

func profitBeforeAndAfterTax(revenue, expences, taxRate float64) (float64, float64) {
	profitBeforeTax := revenue - expences
	profitAfterTax := profitBeforeTax * (1 - taxRate/100)
	return profitBeforeTax, profitAfterTax
}
