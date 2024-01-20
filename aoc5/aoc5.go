package main

import (
	"aoc2023/utils"
	"bufio"
	"os"
)

func main() {
	var data, err = os.Open("aoc5example")
	utils.Check(err)

	var scanner = bufio.NewScanner(data)
	for scanner.Scan() {
		var text = scanner.Text()
		print(text)
	}
}
