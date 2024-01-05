package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) String() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreatedAt.String())
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(strings.ToLower(n.Title), " ", "_") + ".json"
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
