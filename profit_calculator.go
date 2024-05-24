package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Enter revenue: ");
	expenses := getUserInput("Enter expenses: ");
	taxRate := getUserInput("Enter tax rate: ");

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate);

fmt.Printf("Earnings before tax: %.2f\n", ebt);
fmt.Printf("Profit: %.2f\n", profit);
fmt.Printf("Ratio: %.2f\n", ratio);

os.WriteFile("output.txt", []byte(fmt.Sprintf("Earnings before tax: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)), 0644);
}

func getUserInput(message string) float64 {
	var input float64;
	
	fmt.Print(message);
	fmt.Scan(&input);

	switch true {
		case input == 0:
			panic("Input cannot be 0")
		case input < 0:
			panic("Input cannot be negative")
		default:
			break;
	}

	return input;
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses;
	profit := ebt * (1 - (taxRate / 100));
	ratio := ebt / profit;
	return ebt, profit, ratio;
}