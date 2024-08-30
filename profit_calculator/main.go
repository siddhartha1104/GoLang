package main

import "fmt"

func main() {
	var revenue float64
	var expences float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expences: ")
	fmt.Scan(&expences)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	var profitBeforeTax float64
	var profitAfterTax float64

	profitBeforeTax = revenue - expences
	profitAfterTax = profitBeforeTax * (1 - taxRate/100)
	ratio := profitBeforeTax / profitAfterTax

	fmt.Println(profitBeforeTax)
	fmt.Println(profitAfterTax)
	fmt.Println(ratio)
}
