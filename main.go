package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank Account Manager")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var depositAmount float64
		fmt.Print("Your deposit amount: ")
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount
		fmt.Println("Your updated account balance:", accountBalance)
	}
}
