package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse value to float")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
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
			writeBalanceToFile(*accountBalancePtr)
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
			writeBalanceToFile(*accountBalancePtr)
		case 3:
			fmt.Println("Your Account Balance:", *accountBalancePtr)
		default:
			fmt.Print("Goodbye!")
			return
		}
	}

}
