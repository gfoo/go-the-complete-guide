package ops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, floatValue float64) {
	os.WriteFile(fileName, []byte(fmt.Sprint(floatValue)), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("failed to read file (default set to 0)")
	}

	floatValue, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("content of file is not a number: '" + string(data) + "' (default set to 0)")
	}

	return floatValue, nil
}
