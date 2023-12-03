package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

func Part2() {
	start := time.Now()

	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)

	numberPositions := []Position{}
	symbolPositions := []Position{}
	lines := []string{}
	lineIndex := 0
	for fs.Scan() {
		line := fs.Text()
		lines = append(lines, line)

		number := ""
		for i, v := range line {
			if string(v) >= "0" && string(v) <= "9" {
				number += string(v)
			}

			if number != "" && (i == len(line)-1 || string(v) == "." || CheckIsSymbol(string(v))) {
				if string(v) == "." || CheckIsSymbol(string(v)) {
					numberPositions = append(numberPositions, Position{Line: lineIndex, StartIndex: (i - 1) - (len(number) - 1), EndIndex: i - 1, Value: number})
				} else {
					numberPositions = append(numberPositions, Position{Line: lineIndex, StartIndex: i - (len(number) - 1), EndIndex: i, Value: number})
				}
				number = ""
			}

			if string(v) == "*" {
				symbolPositions = append(symbolPositions, Position{Line: lineIndex, StartIndex: i})
			}
		}

		lineIndex++
	}

	gearNumbers := []int{}
	for _, v := range symbolPositions {
		possibleNumbers := []int{}

		if v.StartIndex != 0 {
			if checkIsNumber(string(lines[v.Line][v.StartIndex-1])) {
				number := findNumber(numberPositions, v.Line, v.StartIndex-1, "end")
				possibleNumbers = append(possibleNumbers, number)
			}
		}

		if v.EndIndex != len(lines[v.Line])-1 {
			if checkIsNumber(string(lines[v.Line][v.StartIndex+1])) {
				number := findNumber(numberPositions, v.Line, v.StartIndex+1, "start")
				possibleNumbers = append(possibleNumbers, number)
			}
		}

		if v.Line != 0 {
			numbers := verticalNumberCheck(numberPositions, lines[v.Line-1], v.Line-1, v.StartIndex)
			for _, n := range numbers {
				possibleNumbers = append(possibleNumbers, n)
			}
		}

		if v.Line != len(lines)-1 {
			numbers := verticalNumberCheck(numberPositions, lines[v.Line+1], v.Line+1, v.StartIndex)
			for _, n := range numbers {
				possibleNumbers = append(possibleNumbers, n)
			}
		}

		if len(possibleNumbers) == 2 {
			sum := possibleNumbers[0] * possibleNumbers[1]
			gearNumbers = append(gearNumbers, sum)
		}
	}

	sum := 0
	for _, v := range gearNumbers {
		sum += v
	}

	elpased := time.Since(start)
	fmt.Println(sum, elpased)
}

func checkIsNumber(str string) bool {
	if str >= "0" && str <= "9" {
		return true
	}

	return false
}

func findNumber(numbers []Position, line int, posIndex int, from string) int {
	number := "0"
	if from == "start" {
		for _, n := range numbers {
			if n.Line == line && n.StartIndex == posIndex {
				number = n.Value
				break
			}
		}
	} else {
		for _, n := range numbers {
			if n.Line == line && n.EndIndex == posIndex {
				number = n.Value
				break
			}
		}
	}

	n, _ := strconv.Atoi(number)
	return n
}

func verticalNumberCheck(numbers []Position, line string, lineNum int, posIndex int) []int {
	lineToCheck := ""

	if posIndex == len(line)-1 {
		lineToCheck = line[posIndex-1 : posIndex+1]
	} else if posIndex == 0 {
		lineToCheck = line[posIndex : posIndex+2]
	} else {
		lineToCheck = line[posIndex-1 : posIndex+2]
	}

	possibleNumbers := []int{}
	for i, v := range lineToCheck {
		if i == len(lineToCheck)-1 && checkIsNumber(string(v)) {
			num := findNumber(numbers, lineNum, posIndex+1, "start")
			if num == 0 {
				continue
			}

			possibleNumbers = append(possibleNumbers, num)
			break
		}

		if checkIsNumber(string(v)) && (lineToCheck[i+1]) == '.' && i+1 != len(lineToCheck)-1 {
			num := findNumber(numbers, lineNum, posIndex-1+i, "end")
			if num == 0 {
				continue
			}

			possibleNumbers = append(possibleNumbers, num)
			continue
		}

		if i == 0 && checkIsNumber(string(v)) {
			num := findNumber(numbers, lineNum, posIndex-1, "start")
			if num == 0 {
				continue
			}

			possibleNumbers = append(possibleNumbers, num)
			continue
		}

		if i == len(lineToCheck)-2 && checkIsNumber(string(v)) && (lineToCheck[i+1]) == '.' {
			num := findNumber(numbers, lineNum, posIndex, "end")
			if num == 0 {
				break
			}
			if slices.Contains(possibleNumbers, num) {
				break
			}

			possibleNumbers = append(possibleNumbers, num)
			break
		}

		if checkIsNumber(string(v)) {
			num := findNumber(numbers, lineNum, posIndex, "start")
			if num == 0 {
				break
			}

			possibleNumbers = append(possibleNumbers, num)
			break
		}

	}

	return possibleNumbers
}
