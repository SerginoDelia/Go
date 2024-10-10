// Working with variables, values and operators

package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
	var investmentAmount, expectedReturnRate, years float64
  
  fmt.Print("Investment Amount: ")
  // fmt.Scan get input from the user
  // assign user input to a variable &investmentAmount
  fmt.Scan(&investmentAmount)

  // Get the expected rate of return from the user
  fmt.Print("Expected Return Rate: ")
  fmt.Scan(&expectedReturnRate)

  // Get the number of years from the user
  fmt.Print("Years: ")
  fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

  futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) 
  fmt.Println(futureValue)
  fmt.Println(futureRealValue)
}
