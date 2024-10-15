// Control structures

package main

// to write data to a file we need the "os" package
// strconv to convert strings
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// read data from file
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	// if error, write the code that should run
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	// return balance and nil if we have no errors
	return balance, nil
}

// writeBalanceToFile writes the account balance to a file.
func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		return errors.New("failed to write balance to file")
	}
	return nil
}

// presentOptions displays the menu options to the user.
func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
	} else {
		fmt.Printf("Current balance: %.2f\n", accountBalance)
	}

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
