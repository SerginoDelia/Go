// Control structures

package main

// to write data to a file we need the "os" package
// strconv to convert strings
import (
	"fmt"
	"os"
  "strconv"
)

const accountBalanceFile = "balance.txt"

// read data from file
func getBalanceFromFile() float64 {
  data, _ := os.ReadFile(accountBalanceFile)
  balanceText := string(data)
  balance, _ := strconv.ParseFloat(balanceText, 64)
  return balance
}

// function to handle writing to a file
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// create a balance.txt file, and turn the string to a byte collection, add permissions to the file
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {

		case 1:
			fmt.Println("Your balance is: ", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater that 0.")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if withdrawalAmount >= accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			// We have to use return to end the entire application
			return
			//break
		}
	}
}
