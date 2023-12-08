package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	instruction := ""
	idx := 0
	nodes := make(map[string][]string)
	for fs.Scan() {
		line := fs.Text()

		if idx == 0 {
			instruction = line
		}

		if idx != 0 && line != "" {
			str := strings.Split(line, " = ")
			node := str[0]
			nodeInstrucion := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str[1], "(", ""), ")", ""), ", ")
			nodes[node] = append(nodes[node], nodeInstrucion...)
		}

		idx++
	}

	i := 0
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		direction := nodes[currentNode]

		if instruction[i] == 'L' {
			currentNode = direction[0]
		}

		if instruction[i] == 'R' {
			currentNode = direction[1]
		}

		steps++
		if i == len(instruction)-1 {
			i = 0
		} else {
			i++
		}
	}

	fmt.Println(steps)
}
