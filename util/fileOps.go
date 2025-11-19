package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string){
	textValue := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(textValue), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err :=os.ReadFile(fileName)
	if err != nil{
		return defaultValue, errors.New("failed to find file")
	}
	textValue := string(data)
	value, err := strconv.ParseFloat(textValue, 64)
	if err != nil{
		return defaultValue, errors.New("failed to parse stored value")
	}
	return value, nil
}
