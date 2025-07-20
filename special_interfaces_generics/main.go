package main

import "fmt"

func main() {
	result := addTwoValues(1, 2)

	fmt.Println(result)
}

/*
// ANY VALUE ALLOWED SPECIAL TYPE

func printAnyKindValue(value interface{}) {
	fmt.Println(value)
}

OR USE THE BELOW WAY TO USE SPECIAL TYPE

func printAnyKindValue(value any){
	intVal, ok := value.(int)
	if ok{
		fmt.Println("Integer Value:",intVal)
		return 
	}

	floatVal, ok := value.(float64)
	if ok{
		fmt.Println("Float Value:",floatVal)
		return 
	}

	// so on...
}



// this is thee limitation of Interfaces we neeed to code a lot
func addTwoValues(a, b any) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}
	return nil
}

*/

// Making use of Generics
// Mainly used for libraries etc like udk what it might be use for by users Not commonly used this method

func addTwoValues[T int | float64 | string](a, b T) T {
	return a + b
}

