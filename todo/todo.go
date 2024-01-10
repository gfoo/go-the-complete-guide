package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) String() string {
	return fmt.Sprintf("Title: %s", t.Text)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("text cannot be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (t Todo) Save() error {
	filename := "todo.json"
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
