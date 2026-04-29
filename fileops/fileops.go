package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileNmae string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileNmae, []byte(valueText), 0644)
}

func GetFloatFromFile(fileNmae string) (float64, error) {
	data, err := os.ReadFile(fileNmae)

	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parce stored value.")
	}

	return value, nil
}
