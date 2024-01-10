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

type displayer interface {
	String() string
}

type outputable interface {
	saver     // interface with one methode called save
	displayer // interface with one methode called String
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

func outputData(data outputable) error {
	err := printSomething(data)
	if err != nil {
		return err
	}
	return saveData(data)
}

func printSomething(data interface{}) error {

	note, ok := data.(*note.Note)
	if ok {
		fmt.Println("Note:", note)
		return nil
	}
	todo, ok := data.(*todo.Todo)
	if ok {
		fmt.Println("Todo:", todo)
		return nil
	}
	return fmt.Errorf("do not print content of this type: %T", data)
}

func main() {

	// err := printSomething("Hello")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	title, content := getNoteData()
	text := getTodoData()

	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	err = outputData(userTodo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}
	err = outputData(userNote)
	if err != nil {
		return
	}
}
