package main

import (
	"fmt"

	"example.com/bankaccountmanager-app/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
	}
	accountBalancePtr := &accountBalance
	fmt.Println("Welcome to Go Bank Account Manager")

	for /*i := 0; i < 2; i++ */ {
		operations()

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
				//return
				continue
			}

			*accountBalancePtr += depositAmount
			fmt.Println("Your updated account balance:", *accountBalancePtr)
			fileops.WriteFloatToFile(accountBalanceFile, *accountBalancePtr)
		case 2:
			var withdrawalAmount float64
			fmt.Print("Your Withdrawal Amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > *accountBalancePtr {
				fmt.Println("Invalid amount! Not enough money in your bank account.")
				//return
				continue
			}
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! Should be greater than 0.")
				//return
				continue
			}

			*accountBalancePtr -= withdrawalAmount
			fmt.Println("Your updated account balance:", *accountBalancePtr)
			fileops.WriteFloatToFile(accountBalanceFile, *accountBalancePtr)
		case 3:
			fmt.Println("Your Account Balance:", *accountBalancePtr)
		default:
			fmt.Print("Goodbye!")
			return
		}
	}

}
