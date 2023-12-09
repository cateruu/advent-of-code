package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	allValues := []int{}
	for fs.Scan() {
		line := strings.Split(fs.Text(), " ")

		history := make([]int, len(line))
		for i, v := range line {
			history[i], _ = strconv.Atoi(v)
		}
		allValues = append(allValues, getExtrapolatedValue(history))
	}

	sum := 0
	for _, v := range allValues {
		sum += v
	}

	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
}

func getExtrapolatedValue(history []int) int {
	values := [][]int{}
	values = append(values, history)
	idx := 0
	for !checkAllZeros(values[len(values)-1]) {
		temp := []int{}
		for i, num := range values[idx] {
			if i == 0 {
				continue
			}

			temp = append(temp, (num - values[idx][i-1]))
		}
		values = append(values, temp)
		idx++
	}

	lastNumbers := []int{}
	for _, his := range values {
		lastNumbers = append(lastNumbers, his[len(his)-1])
	}

	return getNextValue(lastNumbers)
}

func checkAllZeros(history []int) bool {
	badNums := 0
	for _, v := range history {
		if v != 0 {
			badNums++
		}
	}

	return badNums == 0
}

func getNextValue(values []int) int {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}

	nextValue := 0
	for i, v := range values {
		if i == 0 {
			continue
		}

		nextValue += v
	}

	return nextValue
}
