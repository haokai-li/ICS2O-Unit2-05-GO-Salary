// Created by: haokai
// Created on: May 2021
// This program displays, "Salary"

package main

import (
	"fmt"
	"github.com/leekchan/accounting"
)

func main() {
	// This function displays currency
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	
	// This function does addition
	var hours float64
	var rate float64
	var your float64
	var govenment float64
	
	// input
	fmt.Println("This program is about salary program.")
	fmt.Println()
	fmt.Print("Enter the number of hours worked: ")
	fmt.Scanln(&hours)
	fmt.Print("Enter the hourly rate: ")
	fmt.Scanln(&rate)
	fmt.Println()
	
	// process
	your = (hours * rate) * (1.00 - 0.18)
	govenment = (hours * rate) * 0.18
	
	// output
	fmt.Println("Your pay will be:", accountingFormater.FormatMoney(your))
	fmt.Println("The govenment will take:", accountingFormater.FormatMoney(govenment))
	fmt.Println("\nDone.")
}
