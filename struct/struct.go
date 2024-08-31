package main

import "fmt"

func main() {

	age := getUserInput("Enter your age: ")
	fmt.Print("age: ", age)

}

func getUserInput(txt string) float64 {
	fmt.Println(txt)
	var value float64
	fmt.Scan(&value)
	return value
}

// todo..
