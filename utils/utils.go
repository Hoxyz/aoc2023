package utils

import (
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseIntsFromString(s string, separator string) []int {
	var numbers []int
	splitString := strings.Split(s, separator)
	for i := 0; i < len(splitString); i++ {
		if splitString[i] != "" {
			number, err := strconv.Atoi(splitString[i])
			Check(err)
			numbers = append(numbers, number)
		}
	}
	return numbers
}
