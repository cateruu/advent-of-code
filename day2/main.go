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

	sum := 0
	for fs.Scan() {
		game := fs.Text()

		rounds := strings.Split(game, ":")[1]
		maxCubes := map[string]int{"red": 1, "green": 1, "blue": 1}

		for _, round := range strings.Split(rounds, ";") {
			for _, cubes := range strings.Split(round, ",") {
				cubeInfo := strings.Split(cubes, " ")
				cubeAmount, _ := strconv.Atoi(cubeInfo[1])
				if maxCubes[cubeInfo[2]] < cubeAmount {
					maxCubes[cubeInfo[2]] = cubeAmount
				}
			}
		}

		setPower := 1
		for _, v := range maxCubes {
			setPower *= v
		}

		sum += setPower
	}

	fmt.Println(sum)
}
