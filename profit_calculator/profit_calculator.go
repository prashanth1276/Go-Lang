package main

import (
	"fmt"
	"errors"
	"os"
)

const fileName = "results.txt"

func writeToFile(rev, exp, tax float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", rev, exp, tax)
	os.WriteFile(fileName, []byte(results), 0644)
}

func readFromFile() (float64, float64, float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, 0, 0, errors.New("file not found")
	}

	var ebt, profit, ratio float64
	_, err = fmt.Sscanf(string(data), "EBT: %f\nProfit: %f\nRatio: %f", &ebt, &profit, &ratio)
	if err != nil {
		return 0, 0, 0, errors.New("failed to parse file data")
	}

	return ebt, profit, ratio, nil
}


func main(){
	revenue, err := getUserInput("Revenue: ")
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeToFile(earningsBeforeTax, profit, ratio)

	ebt, pro, r,err := readFromFile()
	if err != nil {
		fmt.Println("ERROR")
		panic(err)
	}
	fmt.Println("Earnings Before Tax:", ebt)
	fmt.Println("Profit:", pro)
	fmt.Println("Ratio:", r)
}

func getUserInput(infoText string) (float64, error){
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(& userInput)

	if userInput <= 0 {
		return 0, errors.New("it must be positive number")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64){
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate / 100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}