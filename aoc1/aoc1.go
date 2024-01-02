package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetWrittenDigitAsSubstring(str string) int {
	var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for index, digit := range digits {
		if s.Contains(str, digit) {
			return index + 1
		}
	}
	return 0
}

func FindFirstNumberInString(str string) int {
	for i := 0; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) {
			return int(str[i] - '0')
		}
		writtenDigit := GetWrittenDigitAsSubstring(str[:i+1])
		if writtenDigit > 0 {
			return writtenDigit
		}
	}
	return 0
}

func FindLastNumberInString(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(str[i])) {
			return int(str[i] - '0')
		}
		writtenDigit := GetWrittenDigitAsSubstring(str[i:])
		if writtenDigit > 0 {
			return writtenDigit
		}
	}
	return 0
}

func main() {
	var data, err = os.Open("aoc1input")
	check(err)
	var scanner = bufio.NewScanner(data)
	var sum = 0
	for scanner.Scan() {
		var text = scanner.Text()
		var a, b = FindFirstNumberInString(text), FindLastNumberInString(text)
		sum += a*10 + b
	}
	fmt.Println(sum)
}
