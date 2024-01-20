package main

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type card struct {
	winningNumbers []int
	actualNumbers  []int
}

func main() {
	var data, err = os.Open("aoc4input")
	utils.Check(err)

	var cards []card
	var copies []int

	var scanner = bufio.NewScanner(data)
	for scanner.Scan() {
		var text = scanner.Text()
		allNumbers := strings.Split(text, ":")[1]
		winningNumbers := utils.ParseIntsFromString(strings.Split(allNumbers, "|")[0], " ")
		actualNumbers := utils.ParseIntsFromString(strings.Split(allNumbers, "|")[1], " ")

		copies = append(copies, 0)
		cards = append(cards, card{winningNumbers: winningNumbers, actualNumbers: actualNumbers})
	}

	totalCopies := CalculateTotalCopies(cards, copies)
	fmt.Println(totalCopies)
}

func CalculateTotalCopies(cards []card, copies []int) int {
	currentCopies := 1
	totalCopies := 0
	for i := 0; i < len(cards); i++ {
		hits := GetHitsFromCard(cards[i])
		currentCopies += copies[i]
		totalCopies += currentCopies
		if i+1 < len(cards) {
			copies[i+1] += currentCopies
		}
		if i+hits+1 < len(cards) {
			copies[i+hits+1] -= currentCopies
		}
	}
	return totalCopies
}

func GetHitsFromCard(c card) int {
	hits := 0
	for _, number := range c.actualNumbers {
		for _, winningNumber := range c.winningNumbers {
			if number == winningNumber {
				hits++
				break
			}
		}
	}
	return hits
}
