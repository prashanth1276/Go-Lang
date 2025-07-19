package main

import "fmt"

func main(){
	age := 20 // regular variable

	agePointer := & age // var agePointer *int and &age is the age address

	fmt.Println("Age", *agePointer)

	// adultYears := getAdultYears(agePointer)
	// fmt.Println("Adult since:", adultYears, "Years")
	editAgeToAdultYears(agePointer)
	fmt.Println("Age:", age)
}

/*
func getAdultYears(age *int) int {
	return *age - 18
}
*/

// Data Mutation
func editAgeToAdultYears(age *int)  { 
	*age = *age - 18
}