package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

const (
	RED   int = 12
	GREEN int = 13
	BLUE  int = 14
)

func main() {
	var data, _ = os.Open("aoc2input")
	var scanner = bufio.NewScanner(data)
	var answer = 0
	for scanner.Scan() {
		var game = scanner.Text()
		var _ = GetGameId(game)

		var draws = GetGameDraws(game)
		var splitDraws = s.Split(draws, ";")
		var colorDrawList []string
		for i := 0; i < len(splitDraws); i++ {
			var colorDraws = s.Split(splitDraws[i], ",")
			colorDrawList = append(colorDrawList, colorDraws...)
		}

		var maxRed, maxGreen, maxBlue = GetMaxAmountOfAllColors(colorDrawList)

		answer += maxRed * maxGreen * maxBlue
	}
	fmt.Println(answer)
}

func IsValidDraw(maxRed int, maxGreen int, maxBlue int) bool {
	return maxRed <= RED && maxGreen <= GREEN && maxBlue <= BLUE
}

func GetMaxAmountOfAllColors(colorDrawList []string) (int, int, int) {
	var maxRed, maxGreen, maxBlue = 0, 0, 0
	for i := 0; i < len(colorDrawList); i++ {
		var colorDraw = colorDrawList[i]

		var redAmount = GetAmountOfColor(colorDraw, "red")
		var greenAmount = GetAmountOfColor(colorDraw, "green")
		var blueAmount = GetAmountOfColor(colorDraw, "blue")
		if redAmount > maxRed {
			maxRed = redAmount
		}
		if greenAmount > maxGreen {
			maxGreen = greenAmount
		}
		if blueAmount > maxBlue {
			maxBlue = blueAmount
		}
	}

	return maxRed, maxGreen, maxBlue
}

func GetAmountOfColor(colorDraw string, color string) int {
	if s.Contains(colorDraw, color) {
		var amount, err = strconv.Atoi(s.Split(colorDraw, " ")[1])
		checkError(err)
		return amount
	}
	return 0
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetGameDraws(game string) string {
	// Expects input as stated in question ("Game X: ...")
	return s.Split(game, ":")[1]
}

func GetGameId(game string) int {
	var splitGame = s.Split(game, " ")

	// Expects input as stated in question ("Game X: ...")
	var trimmedGameId = s.TrimRight(splitGame[1], ":")
	var id, err = strconv.Atoi(trimmedGameId)
	checkError(err)
	return id
}
