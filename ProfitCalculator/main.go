package main

import (
	"fmt"
)

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Enter Revenue - ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses - ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate (in percentage) - ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses

	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))

	ratio := earningsBeforeTax / profit

	fmt.Print("EBT - ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Profit - ")
	fmt.Println(profit)

	fmt.Print("Ratio - ")
	fmt.Println(ratio)

}
