package fileops


import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil

}

func WriteFloatToFile(value float64, fileName string) {
	valuetext := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuetext), 0644)
}