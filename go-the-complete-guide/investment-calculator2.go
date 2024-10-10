// Working with variables, values and operators

package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
  futureRealValue := futureValue / math.Pow(1+inflationRate/100, years) 
  fmt.Println(futureValue)
  fmt.Println(futureRealValue)
}
