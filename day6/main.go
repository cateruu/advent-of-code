package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	start := time.Now()
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	race := Race{}
	idx := 0
	for fs.Scan() {
		line := strings.Split(fs.Text(), ":")
		nums := strings.TrimSpace(line[1])

		nums = strings.ReplaceAll(nums, " ", "")
		nInt, _ := strconv.Atoi(nums)
		if idx == 0 {
			race = Race{Time: nInt}
		} else {
			race = Race{Time: race.Time, Distance: nInt}
		}
		idx++
	}

	waysToWin := 0
	distance := 0
	for i := 1; i < race.Time; i++ {
		distance = i * (race.Time - i)
		if distance > race.Distance {
			waysToWin++
		}
	}

	elapsed := time.Since(start)
	fmt.Println(waysToWin, elapsed)
}
