package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

type bounds struct {
	line  int
	left  int
	right int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var data, err = os.Open("aoc3input")
	check(err)

	numberMap := make(map[bounds]int)
	var specials [][]int

	var scanner = bufio.NewScanner(data)
	var line = 0
	for scanner.Scan() {
		var text = scanner.Text()
		var specialsOnLine []int
		for i := 0; i < len(text); i++ {
			if unicode.IsDigit(rune(text[i])) {
				var j = i
				for ; j < len(text); j++ {
					if !unicode.IsDigit(rune(text[j])) {
						break
					}
				}
				var number, err = strconv.Atoi(text[i:j])
				check(err)
				numberMap[bounds{line, i, j - 1}] = number
				i = j - 1
			} else if text[i] != '.' {
				specialsOnLine = append(specialsOnLine, i)
			}
		}
		specials = append(specials, specialsOnLine)
		line++
	}
	// PART 1
	//sumOfParts := 0
	//for bound, number := range numberMap {
	//	shouldCount := false
	//	for offset := -1; offset <= 1; offset++ {
	//		if bound.line+offset < 0 || bound.line+offset > len(specials)-1 {
	//			continue
	//		}
	//		for i := 0; i < len(specials[bound.line+offset]); i++ {
	//			special := specials[bound.line+offset][i]
	//			if bound.left-1 <= special && special <= bound.right+1 {
	//				shouldCount = true
	//			}
	//		}
	//	}
	//	if shouldCount {
	//		sumOfParts += number
	//	}
	//}
	//
	//fmt.Println(sumOfParts)
}
