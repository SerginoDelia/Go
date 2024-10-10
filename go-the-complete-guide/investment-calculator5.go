// Working with variables, values and operators

package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, expectedReturnRate, years float64
  
  // fmt.Print("Investment Amount: ")
  outputText("Invest Amount: ")

  // fmt.Scan get input from the user
  // assign user input to a variable &investmentAmount
  fmt.Scan(&investmentAmount)

  // Get the expected rate of return from the user
  // fmt.Print("Expected Return Rate: ")
  outputText("Expected Return Rate: ")
  fmt.Scan(&expectedReturnRate)

  // Get the number of years from the user
  // fmt.Print("Years: ")
  outputText("Years: ")
  fmt.Scan(&years)

	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
  // futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

  futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

  // formattedFV := fmt.Sprintf save the return value to a variable
  formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
  formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

  // fmt.Println("Future value: ", futureValue)
  // fmt.Printf("Future Value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
  // fmt.Println("Future value (adjusted for inflation: ", futureRealValue)
  fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
  fmt.Print(text)
}

// declare return values and types (fv float64, rfv float64)
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
  fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
  rfv = fv / math.Pow(1+inflationRate/100, years)
  return fv, rfv
  // you can just return, because the return values are already declared
  // return
}


















