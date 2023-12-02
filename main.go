package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)

	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	idx := 1
	for fs.Scan() {
		game := fs.Text()

		rounds := strings.Split(game, ":")[1]

		impossible := 0
		for _, round := range strings.Split(rounds, ";") {
			for _, cubes := range strings.Split(round, ",") {
				cubeInfo := strings.Split(cubes, " ")
				cubeAmount, _ := strconv.Atoi(cubeInfo[1])
				if bag[cubeInfo[2]] < cubeAmount {
					impossible++
					break
				}
			}
		}

		if impossible == 0 {
			sum += idx
		}

		idx++
	}

	fmt.Println(sum)
}
