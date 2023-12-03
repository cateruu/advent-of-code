package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	Line       int
	StartIndex int
	EndIndex   int
	Value      string
}

func main() {
	start := time.Now()

	readfile, err := os.Open("input.txt")
	defer readfile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readfile)

	positions := []Position{}
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
				if string(v) == "." {
					positions = append(positions, Position{Line: lineIndex, StartIndex: (i - 1) - (len(number) - 1), EndIndex: i - 1, Value: number})
				} else {
					positions = append(positions, Position{Line: lineIndex, StartIndex: i - (len(number) - 1), EndIndex: i, Value: number})
				}
				number = ""
			}
		}
		lineIndex++
	}

	partNumbers := []int{}
	for _, v := range positions {
		var left, right, top, bottom bool

		if v.StartIndex != 0 {
			left = CheckIsSymbol(string(lines[v.Line][v.StartIndex-1]))
		}

		if v.EndIndex != len(lines[v.Line])-1 {
			right = CheckIsSymbol(string(lines[v.Line][v.EndIndex]))
		}

		if v.Line != 0 {
			top = verticalCheck(lines[v.Line-1], v)
		}

		if v.Line != len(lines)-1 {
			bottom = verticalCheck(lines[v.Line+1], v)
		}

		if left || right || top || bottom {
			number, _ := strconv.Atoi(v.Value)
			partNumbers = append(partNumbers, number)
		}
	}

	sum := 0
	for _, v := range partNumbers {
		sum += v
	}

	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
	Part2()
}

func CheckIsSymbol(str string) bool {
	symbols := `+-*/@&$#=%`

	return strings.Contains(symbols, str)
}

func verticalCheck(line string, numInfo Position) bool {
	lineToCheck := ""

	if numInfo.EndIndex == len(line)-1 {
		lineToCheck = line[numInfo.StartIndex-1 : numInfo.EndIndex+1]
	} else if numInfo.StartIndex == 0 {
		lineToCheck = line[numInfo.StartIndex : numInfo.EndIndex+2]
	} else {
		lineToCheck = line[numInfo.StartIndex-1 : numInfo.EndIndex+2]
	}

	for _, s := range lineToCheck {
		if CheckIsSymbol(string(s)) {
			return true
		}
	}

	return false
}
