package main

import (
	"fmt"
	"math"
)

func main() {

	var loanAmount, annualRate, totalPayments, monthlyRate float64

	// Get user input
	loanAmount = getInfo("Loan Amount: ")
	annualRate = getInfo("Annual Interest Rate (in %): ") / 100 // Convert to decimal
	totalPayments = getInfo("Total Loan Period (Years): ") * 12 // Convert years to months

	// Monthly interest rate
	monthlyRate = annualRate / 12

	// Calculate monthly payment using the correct formula
	monthlyPayment := loanAmount * (monthlyRate * math.Pow(1+monthlyRate, totalPayments)) / (math.Pow(1+monthlyRate, totalPayments) - 1)

	// Calculate total amount and overpayment
	totalAmount := monthlyPayment * totalPayments
	overPayment := totalAmount - loanAmount

	// Output the result
	fmt.Printf("Total amount of payments: %.2f\n", totalAmount)
	fmt.Printf("Overpayment: %.2f\n", overPayment)
}

// Helper function to get user input
func getInfo(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}
