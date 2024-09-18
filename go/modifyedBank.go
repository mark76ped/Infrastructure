package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64, error {
	data, err := os.ReadFile(accountBalanceFile)

		if err != nil {
			return 0, errors.New("Failed to find balance.")
		}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func main() {
	var depositAmount float64
	var accountBalance, err = getBalanceFromFile() // Get balance from the file

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		panic("An Error Occured")
	}
	
	var choice int
	
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Scan(&choice)

		accountBalance = printChoice(choice, accountBalance, depositAmount) // Update balance
		if choice == 4 {
			break
		}
	}
}

func printChoice(choice int, accountBalance float64, depositAmount float64) float64 {
	switch choice {
	case 1:
		fmt.Println("1. Check balance. Your balance is:", accountBalance)
	case 2:
		fmt.Println("2. Deposit Money:")
		fmt.Print("Your desired amount: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Deposited", depositAmount)
		fmt.Println("Your updated balance is", accountBalance)
		writeBalanceToFile(accountBalance) // Write updated balance to file
	case 3:
		fmt.Println("3. Withdraw money")
		fmt.Print("Your desired amount: ")
		fmt.Scan(&depositAmount)
		if depositAmount <= accountBalance {
			accountBalance -= depositAmount
			fmt.Println("Withdrawn", depositAmount)
			fmt.Println("Your updated balance is", accountBalance)
			writeBalanceToFile(accountBalance) // Write updated balance to file
		} else {
			fmt.Println("You do not have enough funds for this transaction.")
		}
	case 4:
		fmt.Println("4. Exit. Goodbye!")
	default:
		fmt.Println("Invalid option, please try again.")
	}
	return accountBalance // Return updated balance
}