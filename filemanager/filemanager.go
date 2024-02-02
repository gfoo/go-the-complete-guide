package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InuptFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InuptFilePath)
	if err != nil {
		return nil, errors.New("Could not open the file " + fm.InuptFilePath)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("Error reading the file " + fm.InuptFilePath)
	}
	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Could not create the file " + fm.OutputFilePath)
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close()
		return errors.New("Could not encode the data as JSON")
	}
	// file.Close()
	return nil
}

func News(inuptFilePath string, outputFilePath string) FileManager {
	return FileManager{
		InuptFilePath:  inuptFilePath,
		OutputFilePath: outputFilePath,
	}
}
