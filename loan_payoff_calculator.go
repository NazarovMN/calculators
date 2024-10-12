package main

import (
	"fmt"
	"math"
)

func main() {
	var monthlyPayment, loanAmount, monthlyRate float64

	monthlyPayment = getInfo("monthly payment :")
	loanAmount = getInfo("amount of loan :")
	monthlyRate = getInfo("monthly loan rate: ") / 100

	amountMonth := math.Log(monthlyPayment/(monthlyPayment-loanAmount*monthlyRate)) / math.Log(1+monthlyRate)
	totalPaid := amountMonth * monthlyPayment
	totalInterst := totalPaid - loanAmount

	fmt.Printf("Total Number of Month: %.1f\n Total Amount of Interest: %.1f\n", amountMonth, totalInterst)

}

func getInfo(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}
