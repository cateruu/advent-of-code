package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)

	sum := 0
	for fs.Scan() {
		game := strings.Split(fs.Text(), ":")[1]

		numbers := strings.Split(game, "|")
		winningNumbers := strings.Split(numbers[0], " ")[1:]
		ourNumbers := strings.Split(numbers[1], " ")[1:]

		won := 0
		for _, n := range ourNumbers {
			if slices.Contains(winningNumbers, n) && n != "" {
				won++
			}
		}

		points := 0
		for i := 1; i <= won; i++ {
			if i == 1 {
				points += 1
			} else {
				points *= 2
			}
		}

		sum += points
	}

	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
}
