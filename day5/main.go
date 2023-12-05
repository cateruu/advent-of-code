package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

	seeds := []int{}
	maps := [][][]int{}

	tempArr := [][]int{}
	idx := -1
	for fs.Scan() {
		text := fs.Text()
		line := strings.Split(text, ": ")

		if line[0] == "" {
			idx++
		}

		if idx != 0 && line[0] == "" {
			maps = append(maps, tempArr)
			tempArr = nil
		}

		if line[0] == "seeds" {
			s := strings.Split(line[1], " ")
			for i, v := range s {
				if i+1 == len(s) {
					break
				}

				if i%2 == 1 {
					continue
				}

				num, _ := strconv.Atoi(v)
				repeat, _ := strconv.Atoi(s[i+1])
				for j := 0; j < repeat; j++ {
					seeds = append(seeds, num+j)
				}
			}
		} else if line[0] != "" && line[0][len(line[0])-1] != ':' {
			nums := strings.Split(line[0], " ")
			numsArr := []int{}
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				numsArr = append(numsArr, num)
			}
			tempArr = append(tempArr, numsArr)
		}
	}
	// add last map to maps slice
	maps = append(maps, tempArr)
	tempArr = nil

	locations := []int{}
	for _, seed := range seeds {
		value := seed
		for i := 0; i < len(maps); i++ {
			for _, row := range maps[i] {
				if value >= row[1] && value <= row[1]+row[2]-1 {
					diff := value - row[1]
					value = row[0] + diff
					break
				}
			}
		}
		locations = append(locations, value)
	}

	elapsed := time.Since(start)
	lowest := slices.Min(locations)
	fmt.Println(lowest, elapsed)
}
