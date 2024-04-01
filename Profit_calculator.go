package main // entry point for the application

import "fmt"

func main() {
	var expenses, revenue, taxRate float64

	fmt.Print("Enter Revenue")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses")
	fmt.Scan(&expenses)

	earningsBeforeTax := revenue - expenses

	fmt.Println(earningsBeforeTax)

	fmt.Print("enter tax rate")
	fmt.Scan(&taxRate)

	earningAfterTax := (revenue - ((taxRate / 100) * revenue)) - expenses

	fmt.Println(earningAfterTax)

	ratio := earningsBeforeTax / earningAfterTax

	fmt.Print(ratio)

}
