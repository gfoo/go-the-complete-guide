package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"org.gfoo/go-project/note"
	"org.gfoo/go-project/todo"
)

// saver because interface with one methode called save
type saver interface {
	Save() error
}

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
	title := getUserInput("Note Title?")
	content := getUserInput("Note Content?")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo Text?")
	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving:", err)
		return err
	}
	fmt.Println("Saved successfully!")
	return nil
}

func main() {
	title, content := getNoteData()
	text := getTodoData()

	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	fmt.Println(userTodo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	fmt.Println(userNote)

	err = saveData(userTodo)
	if err != nil {
		return
	}

	err = saveData(userNote)
	if err != nil {
		return
	}
}
