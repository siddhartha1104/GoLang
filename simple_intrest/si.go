package main

import "fmt"

func main() {

	var principle float64
	var rate float64
	var time float64

	fmt.Print("Enter Principle: ")
	fmt.Scan(&principle)

	fmt.Print("Enter Rate: ")
	fmt.Scan(&rate)

	fmt.Print("Enter Time: ")
	fmt.Scan(&time)

	simpleIntrest := (principle * rate * time)
	print(simpleIntrest)
}
