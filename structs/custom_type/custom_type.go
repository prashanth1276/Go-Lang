package main

import "fmt"

type s string

func (text s) log() {
	fmt.Println(text)
}

func printName(name string) {
	fmt.Println(name)
}

func main() {
	var name s = "Max" // This means now Name is s Type and the u r trying to work on name.log() i.e, name is 'Max' os 'Max'.log() now prints Max thats it
	name.log()

	text := "Prashanth"
	printName(text)
	
}