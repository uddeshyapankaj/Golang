package main // entry point for the application

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount = 100000
	// var expectedReturnRate = 5.5
	// var years = 10

	//type conversion
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// fmt.Println(futureValue)

	// telling explicitly that consider it has float64
	// var investmentAmount float64 = 100000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// fmt.Println(futureValue)

	// shorter syntax
	var investmentAmount float64 = 100000
	expectedReturnRate := 5.5
	years := 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)

}
