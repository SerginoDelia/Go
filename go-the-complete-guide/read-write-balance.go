package main

import (
  "fmt"
  "log"
  "os"
  // converts strings
  "strconv"
)

// Write balance to a file
// create a helper function that write balance to file
func writeToFile(balance float64){

  // convert balance to text
  balanceText := fmt.Sprint(balance)
  // write balance to "balance-test-txt"
  os.WriteFile("balance-test.txt", []byte(balanceText), 0644)
  fmt.Println("Balance successfully written to file")
}

func readFile(file string) {
  readBalance, err := os.ReadFile(file)
  
  // log error if it's not nil
  if err != nil {
    log.Fatal(err)
  }
  // Print the contents of the file
  fmt.Println(string(readBalance))
}

const accountBalanceFile = "balance-test.txt"

func getBalanceFromFile() float64 {
  data, _ := os.ReadFile(accountBalanceFile)
  balanceText := string(data)
  // converts balance to a float64
  balance, _ :=  strconv.ParseFloat(balanceText, 64)
  return balance
}

func main() {
  balance := getBalanceFromFile()
  writeToFile(balance)

  // read file
  readFile("balance-test.txt")
}



















