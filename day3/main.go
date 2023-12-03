package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	lineIndex := 0
	for fs.Scan() {
		line := fs.Text()

		number := ""
		for i, v := range line {
			if number != "" && (string(v) == "." || checkIsSymbol(string(v))) {
				positions = append(positions, Position{Line: lineIndex, StartIndex: (i - 1) - (len(number) - 1), EndIndex: i - 1, Value: number})
				number = ""
			}

			if string(v) >= "0" && string(v) <= "9" {
				number += string(v)
			}

			if checkIsSymbol(string(v)) {
				positions = append(positions, Position{Line: lineIndex, StartIndex: i, Value: string(v)})
			}

		}
		lineIndex++
	}

	symbolPositions := getSymbolPositions(positions)

	partNumbers := []int{}
	for _, v := range positions {
		if checkIsSymbol(v.Value) {
			continue
		}

		previousLine := v.Line - 1
		if v.Line == 0 {
			previousLine = 0
		}
		nextLine := v.Line + 1

		startLookIndex := v.StartIndex - 1
		if v.StartIndex == 0 {
			startLookIndex = 0
		}
		endLookIndex := v.EndIndex + 1
		for i := previousLine; i <= nextLine; i++ {
			for j := startLookIndex; j <= endLookIndex; j++ {
				if isNearSymbol(symbolPositions, i, j) {
					number, _ := strconv.Atoi(v.Value)
					partNumbers = append(partNumbers, number)
				}
			}
		}
	}

	sum := 0
	for _, v := range partNumbers {
		sum += v
	}

	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
}

func checkIsSymbol(str string) bool {
	symbols := []string{"!", "@", "#", "$", "%", "^", "&", "*", "-", "_", "=", "+", "/"}

	return slices.Contains(symbols, str)
}

func getSymbolPositions(positions []Position) []Position {
	symbols := []Position{}
	for _, v := range positions {
		if checkIsSymbol(v.Value) {
			symbols = append(symbols, v)
		}
	}

	return symbols
}

func isNearSymbol(symbols []Position, line int, place int) bool {
	for _, symbol := range symbols {
		if symbol.Line == line && symbol.StartIndex == place {
			return true
		}
	}

	return false
}
