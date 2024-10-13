// Profit Calculator
// Ask For revenue, expenses, and tax rate
// Calculate earnings before tax (EBT) and earnings after tax(profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio

package main

import (
  "fmt"
  "os"
  //"strconv"
)

func writeToFile(msg string, data float64){
  dataText := fmt.Sprint(data)
  os.WriteFile("finances.txt", []byte(msg + dataText), 0644)
}

func main() {
  revenue := getUserInput("Revenue: ") 
  expenses := getUserInput("Expenses: ") 
  taxRate := getUserInput("Tax Rate: ")
  
  ebt, profit, ratio := calculateKpi(revenue, expenses, taxRate)
  fmt.Printf("Earnings before tax: %.2f\n", ebt)
  
  writeToFile("Earnings before tax: $", ebt)
  
  fmt.Printf("Profit: %.2f\n", profit)

  fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
  var userInput float64
  fmt.Print(infoText)
  fmt.Scan(&userInput)
  return userInput
}


//func getUserInput(infoText) (revenue float64, expenses float64, taxRate float64) {
//  fmt.Print("Revenue: ")
//  fmt.Scan(&revenue)
//  fmt.Print("Expenses: ")
//  fmt.Scan(&expenses)
//  fmt.Print("Tax Rate: ")
//  fmt.Scan(&taxRate)

//  return revenue, expenses, taxRate
//}

func calculateKpi(revenue, expenses, taxRate float64) (float64, float64, float64) {
  ebt := revenue - expenses
  profit := ebt * (1 - taxRate/100)
  ratio := ebt/profit
  return ebt, profit, ratio
}









