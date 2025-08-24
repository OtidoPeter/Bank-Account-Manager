package main

import "fmt"

func main() {
	accountBalance := 1000.0
	accountBalancePtr := &accountBalance
	fmt.Println("Welcome to Go Bank Account Manager")

	for /*i := 0; i < 2; i++ */ {
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
			fmt.Print("Your Deposit Amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0.")
				return
			}

			*accountBalancePtr += depositAmount
			fmt.Println("Your updated account balance:", *accountBalancePtr)
		case 2:
			var withdrawalAmount float64
			fmt.Print("Your Withdrawal Amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > *accountBalancePtr {
				fmt.Println("Invalid amount! Not enough money in your bank account.")
				return
			}
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0.")
				return
			}

			*accountBalancePtr -= withdrawalAmount
			fmt.Println("Your updated account balance:", *accountBalancePtr)
		case 3:
			fmt.Println("Your Account Balance:", *accountBalancePtr)
		default:
			fmt.Print("Goodbye!")
		}
	}

}
