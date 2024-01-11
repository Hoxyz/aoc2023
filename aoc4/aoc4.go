package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
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
			check(err)
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func main() {
	var data, err = os.Open("aoc4input")
	check(err)

	answer := 0

	var scanner = bufio.NewScanner(data)
	for scanner.Scan() {
		var text = scanner.Text()
		allNumbers := strings.Split(text, ":")[1]
		winningNumbers := ParseIntsFromString(strings.Split(allNumbers, "|")[0], " ")
		actualNumbers := ParseIntsFromString(strings.Split(allNumbers, "|")[1], " ")

		hits := -1
		for _, number := range actualNumbers {
			for _, winningNumber := range winningNumbers {
				if number == winningNumber {
					hits++
					break
				}
			}
		}

		if hits >= 0 {
			answer += 1 << hits
		}
	}
	fmt.Println(answer)
}
