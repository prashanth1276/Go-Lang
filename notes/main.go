package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface{
	saver
	DisplayData()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	notesApp, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	displayData(todo)

	displayData(notesApp)
}

func displayData(data outputtable) error {
	data.DisplayData()
	return saveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(message string) string {
	// var userInput string
	fmt.Print(message)
	// fmt.Scanln(& userInput)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil{
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data Failed!")
		return err
	}
	fmt.Println("Saving the data succeeded!")
	return nil
}