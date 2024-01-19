package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Could not open the file " + path)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error reading the file " + path)
	}
	file.Close()
	return lines, nil
}

func WriteJSON(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Could not create the file " + path)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Could not encode the data as JSON")
	}
	file.Close()
	return nil
}
