package main

import (
	"fmt"
	"math"
)

func main() {
	var monthlyPayment, loan, monthlyInterestRate, numberMonths float64

	loan = getInfo("Loan Amount: ")
	monthlyInterestRate = getInfo("Monthly Interest Rate (in %): ") / 100 // Convert percentage to decimal
	numberMonths = getInfo("Number of Months: ")

	// Formula to calculate the monthly payment
	monthlyPayment = loan * monthlyInterestRate * math.Pow(1+monthlyInterestRate, numberMonths) /
		(math.Pow(1+monthlyInterestRate, numberMonths) - 1)

	// Output the result
	fmt.Printf("Monthly Payment: %.2f\n", monthlyPayment)
}

// Helper function to get user input
func getInfo(text string) float64 {
	var info float64
	fmt.Print(text)
	fmt.Scan(&info)
	return info
}
