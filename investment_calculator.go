package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	/* 
	This is a Multi line comment
	// -> this is a Single line comment.
	*/
	var years = 10.0
	expectedReturnRate := 5.5
	
	fmt.Println("Enter the Amount You Want to Invest: ")
	var investAmount float64
	fmt.Scan(& investAmount)

	futureValue, futureRealValue := calculateFutureValues(investAmount, expectedReturnRate, years)
	// futureValue := investAmount * math.Pow((1 + expectedReturnRate / 100), years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Using Sprintf Future Value: %.2f \n", futureValue)
	formattedRFV := fmt.Sprintf("Using Sprintf Future Real Value: %.2f \n", futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)
	outputText(formattedFV, formattedRFV) // making use of Custom function
	fmt.Println("Using Println Future Value: ",futureValue)
	fmt.Println("Using Println Future Real Value: ",futureRealValue)
	fmt.Printf("Using Printf Future Value: %v \n", futureValue)
	fmt.Printf("Using Printf Future Real Value: %v \n", futureRealValue)

	// fmt.Println(`formatted
	// 			FRM`)

	fmt.Printf("The sum of Two Numbers are: %f",sumOfTwoNumbers(4.0, 5.0)) // Sumation of Two Numbers Custom Function
	fmt.Printf("\nThe sum of Two Numbers are: %d",multiplyTwoVariables(2, 5)) // Multiplication of Two Numbers
}

func outputText(input1, input2 string){
	fmt.Print(input1, input2)
}

func sumOfTwoNumbers(a, b float64) float64{
	return a + b
}

func calculateFutureValues(iA, eRR, years float64) (float64, float64){
	fv := iA * math.Pow((1 + eRR / 100), years)
	rfv := fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}

/*
ALTERNATIVE WAY TO RETURN IN FUNCTION

func calculateFutureValues(iA, eRR, years float64) (fv float64, rfv float64){  //U only Specify here the variable type and name
	fv = iA * math.Pow((1+eRR/100), years)  //No need to initialize (:=) just assign
	rfv = fv / math.Pow(1+inflationRate/100, years)  //No need to initialize (:=) just assign
	return  //No need to Specify what u want to return as GoLang will take up the Mentioned return type and variable
}
*/

// Making use of Alternative Return type
func multiplyTwoVariables(a, b int) (result int){
	result = a * b
	return
}
