package main

import "fmt"

func main() {

	fmt.Print("\n\t****Welcome to Your Bank****\n\n")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Choose: ")
	fmt.Scan(&choice)

}
