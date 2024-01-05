package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"org.gfoo/go-project/note"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt, " ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Title?")
	content := getUserInput("Content?")
	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
	// Use the note object as needed
}
