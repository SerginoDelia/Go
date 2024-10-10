// Profit Calculator
// Ask For revenue, expenses, and tax rate
// Calculate earnings before tax (EBT) and earnings after tax(profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio

package main

import (
  "fmt"
 // "math"
)

func main() {
  var revenue, expenses, taxRate float64
  // get revenue from user
  fmt.Print("Revenue: ")
  fmt.Scan(&revenue)

  // get expenses
  fmt.Print("Expenses: ")
  fmt.Scan(&expenses)

  // get tax rate
  fmt.Print("Tax Rate: ")
  fmt.Scan(&taxRate)
  
  // earnings before tax = revenue - expenses
  earningsBeforeTax := revenue - expenses
  fmt.Println("Earnings before tax: ", earningsBeforeTax)
  
  // earning after tax = revenue - expenses - taxes
  // profit := earningsBeforeTax - (earningsBeforeTax * taxRate / 100)
  profit := earningsBeforeTax * (1 - taxRate/100)
  fmt.Println("Profit: ", profit)

  // calculate ratio
  ratio := earningsBeforeTax / profit
  fmt.Println("Ratio: ", ratio)
}














