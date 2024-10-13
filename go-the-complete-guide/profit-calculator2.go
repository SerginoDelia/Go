// Profit Calculator
// Ask For revenue, expenses, and tax rate
// Calculate earnings before tax (EBT) and earnings after tax(profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio

package main

import (
  "fmt"
  "os"
  "errors"
)

func main() {
  revenue, err1 := getUserInput("Revenue: ")

  if err1 != nil {
    fmt.Println(err1)
    return
  }

  expenses, err2 := getUserInput("Expenses: ")

  if err2 != nil {
    fmt.Println(err2)
    return
  }

  taxRate, err3 := getUserInput("Tax Rate: ")
 
  if err3 != nil {
    fmt.Println(err3)
    return
  }

  //if err1 != nil || err2 != nil || err3 != nil {
    //fmt.Println(err1)
    //return
  //}  

  ebt, profit, ratio := calculateKpi(revenue, expenses, taxRate)
  fmt.Printf("Earnings before tax: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
  
  //fmt.print("Earnings before tax: $", ebt)
  //fmt.Printf("Profit: %.2f\n", profit)
  //fmt.Printf("Ratio: %.3f\n", ratio)

  storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
  var userInput float64
  fmt.Print(infoText)
  fmt.Scan(&userInput)
  
  if userInput <= 0 {
    return 0, errors.New("Value must be a positive number.")
  }
 
  return userInput, nil
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

func storeResults(ebt, profit, ratio float64){
  financesText := fmt.Sprintf("Earnings before tax %.2f:\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio )
  os.WriteFile("results.txt", []byte(financesText), 0644)
}







