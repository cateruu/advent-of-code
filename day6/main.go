package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	races := make([]Race, 4)
	idx := 0
	for fs.Scan() {
		line := strings.Split(fs.Text(), ":")
		nums := strings.TrimSpace(line[1])

		innerIdx := 0
		for _, n := range strings.Split(nums, " ") {
			if n == "" {
				continue
			}

			number, _ := strconv.Atoi(n)
			if idx == 0 {
				races[innerIdx] = Race{Time: number}
			} else {
				races[innerIdx] = Race{Time: races[innerIdx].Time, Distance: number}
			}

			innerIdx++
		}
		idx++
	}

	sum := 1
	for _, race := range races {
		waysToWin := 0
		distance := 0
		for i := 1; i < race.Time; i++ {
			distance = i * (race.Time - i)
			if distance > race.Distance {
				waysToWin++
			}
		}

		sum *= waysToWin
	}

	fmt.Println(sum)
}
