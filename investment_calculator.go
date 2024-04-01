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
	const inflationRate = 10.5
	var investmentAmount float64 = 100000
	expectedReturnRate := 5.5
	var years float64

	// Scan - does not work for string with multiple words
	// var formulaName string = "hello world"

	// fmt.Print("formaula-Name: ")
	// fmt.Scan(&formulaName)

	// fmt.Println(formulaName)

	//taking input from user
	fmt.Print("Investement Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	newinflatedFutureValue := futureValue * math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(newinflatedFutureValue)

}
