package fileops

import(
	"strconv"
	"os"
	"fmt"
	"errors"
)

var Hello string = "Prashanth"


func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) // 0644 means owner can real and write the file but others can only read it
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}
	return value, nil
}