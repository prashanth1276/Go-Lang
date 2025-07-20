package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type Notes struct {
	NotesTitle string `json:"notes_title"`
	NotesContent string `json:"notes_content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Notes) Save() error {
	fileName := strings.ReplaceAll(n.NotesTitle, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (n Notes) DisplayData(){
	fmt.Printf("Your Note titled %v has the following Content: \n\n%v\n",n.NotesTitle, n.NotesContent)
}

func New(userNotesTitle, userNotesContent string) (Notes, error) {
	if userNotesTitle == "" || userNotesContent == "" {
		return Notes{}, errors.New("invalid input")
	}
	return Notes{
		NotesTitle: userNotesTitle,
		NotesContent: userNotesContent,
		CreatedAt: time.Now(),
	}, nil
}