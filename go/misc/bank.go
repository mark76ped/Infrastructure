package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
var depositAmount float64
var accountBalance = getBalanceFromFile()
var choice int
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)

	printChoice(choice, accountBalance, depositAmount)
}
func printChoice(choice int, accountBalance float64, depositAmount float64) int {
	
	for {
		switch choice {
			case 1:
				fmt.Println("1. Check balance. Your balance is:", accountBalance)
				fmt.Scan(&choice)
				continue
			case 2:
				fmt.Println("2. Deposit Money:")
				fmt.Print("Your desired amount: ")
				fmt.Scan(&depositAmount)
				fmt.Println("Deposited", depositAmount)
				accountBalance = accountBalance + depositAmount
				fmt.Println("Your updated balance is", accountBalance)
				fmt.Scan(&choice)
				writeBalanceToFile(accountBalance)
				continue
			case 3:
				fmt.Println("3. Withdraw money")
				fmt.Print("Your desired amount: ")
				fmt.Scan(&depositAmount)
				fmt.Println("Withdraw", depositAmount)
					if depositAmount < accountBalance { 
						accountBalance = accountBalance - depositAmount
						fmt.Println("Your updated balance is", accountBalance)
					} else {
						fmt.Println("You do not have enough funds for this transaction, goodbye.")
					}
					fmt.Scan(&choice)
				writeBalanceToFile(accountBalance)
				continue
			default:
					fmt.Println("4. Exit. Goodbye")
				break
		}
	// if choice == 1 {
	// 	fmt.Println("1. Check balance. Your balance is:", accountBalance)
	// 	fmt.Scan(&choice)
	// 	continue
	// 	} else if  choice == 2 {
	// 	fmt.Println("2. Deposit Money:")
	// 	fmt.Print("Your desired amount: ")
	// 	fmt.Scan(&depositAmount)
	// 	fmt.Println("Deposited", depositAmount)
	// 	accountBalance = accountBalance + depositAmount
	// 	fmt.Println("Your updated balance is", accountBalance)
	// 	fmt.Scan(&choice)
	// 	continue
	// 	} else if  choice == 3 {
	// 	fmt.Println("3. Withdraw money")
	// 	fmt.Print("Your desired amount: ")
	// 	fmt.Scan(&depositAmount)
	// 	fmt.Println("Withdraw", depositAmount)
	// 		if depositAmount < accountBalance { 
	// 			accountBalance = accountBalance - depositAmount
	// 			fmt.Println("Your updated balance is", accountBalance)
	// 		} else {
	// 			fmt.Println("You do not have enough funds for this transaction, goodbye.")
	// 		}
	// 		fmt.Scan(&choice)
	// 		continue
	// 	} else {
	// 	fmt.Println("4. Exit. Goodbye")
	// 	break
	// 	}
	}
	return choice
}
