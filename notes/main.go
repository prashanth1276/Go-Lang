package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()

	notesApp, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	notesApp.GetNotesData()
	err = notesApp.Save()

	if err != nil {
		fmt.Println("Saving the Note Failed!")
		return
	}

	fmt.Println("Saving the Note succeeded!")
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