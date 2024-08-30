package main

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalanceTxt string = "accountBalance.txt"

func getBalanceAmount() float64 {
	data, _ := os.ReadFile(accountBalanceTxt)
	balanceTxt := string(data)
	balance, _ := strconv.ParseFloat(balanceTxt, 64)
	return balance
}

func storeBalanceAmount(balance float64) {
	//0644 -> helps to take permissions from the system
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceTxt, []byte(balanceText), 0644)
}

func main() {

	var balance = getBalanceAmount()

	fmt.Print("\n\t****Welcome to Your Bank****\n\n")

	//we can also use switch insted of else if

	for {

		fmt.Println("\n1. Check Balance")
		fmt.Println("2. Deposite Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your Balance is: %v\n", balance)
			storeBalanceAmount(balance)

		} else if choice == 2 {
			var depostieAmount float64
			fmt.Println("Enter the amount: ")
			fmt.Scan(&depostieAmount)

			if depostieAmount <= 0 {
				fmt.Println("\tEnter a valid amount!")
				continue
			}

			balance += depostieAmount
			fmt.Println("Total amount is: ", balance)
			storeBalanceAmount(balance)

		} else if choice == 3 {
			var withdraw float64
			fmt.Println("\tTotal Balance in Your Account: ", balance)
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("\tEnter a valid amount!")
				continue
			}

			if withdraw > balance {
				fmt.Println("\tYour Account balance is less than withdraw amount!")
				continue
			}

			balance -= withdraw
			fmt.Printf("Remaining Balance: %v\n", balance)
			storeBalanceAmount(balance)

		} else {
			fmt.Println("\t!!Thank You !!")
			break

		}
	}

}
