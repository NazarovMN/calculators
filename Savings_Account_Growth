package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit, annualRate, interestPerYear, years float64

	deposit = getInfo("Deposit: ")
	annualRate = getInfo("Annual Rate (in %) :") / 100
	interestPerYear = getInfo("Interest Copmound Per Year:")
	years = getInfo("How many years: ")

	futureValue := deposit * math.Pow(1+annualRate/interestPerYear, interestPerYear*years)
	interesEarned := futureValue - deposit

	fmt.Printf("Total Amount: %.1f\n Earned Percent: %.1f\n", futureValue, interesEarned)

}

func getInfo(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}
