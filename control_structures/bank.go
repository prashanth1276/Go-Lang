package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main(){
	var userChoice int
	fmt.Println(randomdata.PhoneNumber()) // Making use of Third part Package
	balance, err := fileops.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		panic("Can't continue, Sorry!")
	}

	bankLogic(userChoice, balance)
}

func bankLogic(userChoice int, balance float64) {
	fmt.Println("WELCOME TO GO BANK!")
	for {
		presentOptions()
		fmt.Print("Your choice? ")
		fmt.Scan(& userChoice)

		/*
		if userChoice == 1 {
			fmt.Println("Balance:",balance)
		} else if userChoice == 2 {
			deposit := userInput("Deposit Ammount? ")
			if deposit <= 0 {
				fmt.Println("Less than 0 not allowed.")
			} else {
				balance += deposit
				fmt.Println("Successfully Depositted!")
			}
		} else if userChoice == 3 {
			userWithdraws:=userInput("How much you want to Withdraw? ")
			if userWithdraws > balance{
				fmt.Println("Insufficiant Balance!")
				} else {
					balance -= userWithdraws
					fmt.Println("Successfully withdrawn!")
				}
		} else {
			fmt.Println("THANK YOU FOR USING GO BANK! HAVE A GREAT DAY")
			break
		}
		*/

		switch userChoice {
			case 1:
				fmt.Println("Balance:",balance)
			case 2:
				deposit := userInput("Deposit Ammount? ")
				if deposit <= 0 {
					fmt.Println("Less than 0 not allowed.")
				} else {
					balance += deposit
					fileops.WriteFloatToFile(accountBalanceFile, balance)
					fmt.Println("Successfully Depositted!")
				}
			case 3:
				userWithdraws:=userInput("How much you want to Withdraw? ")
				if userWithdraws > balance{
					fmt.Println("Insufficiant Balance!")
					} else {
						balance -= userWithdraws
						fileops.WriteFloatToFile(accountBalanceFile, balance)
						fmt.Println("Successfully withdrawn!")
					}
			default:
				fmt.Println("THANK YOU FOR USING GO BANK! HAVE A GREAT DAY")
				return
		}
	}
}

func userInput(messageToDisplay string) float64 {
	var userInput float64
	fmt.Print(messageToDisplay)
	fmt.Scan(& userInput)
	return userInput
}