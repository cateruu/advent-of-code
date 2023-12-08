package main

import (
	"bufio"
	"fmt"
	"os"
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
	currentNodes := findStartingNodes(&nodes)
	steps := []int{}
	for i := 0; i < len(currentNodes); i++ {
		steps = append(steps, getNodeSteps(currentNodes[i], nodes, instruction))

	}

	elapsed := time.Since(start)
	fmt.Println(lcm(steps...), elapsed)
}

func findStartingNodes(nodes *map[string][]string) []string {
	startingNodes := []string{}
	for k := range *nodes {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}

	return startingNodes
}

func checkIfEnd(nodes *[]string) bool {
	endNodes := 0
	for _, node := range *nodes {
		if node[2] == 'Z' {
			endNodes++
		}
	}

	return endNodes == len(*nodes)
}

func getNodeSteps(startingNode string, nodes map[string][]string, instruction string) int {
	i := 0
	currentNode := startingNode
	steps := 0
	for !strings.HasSuffix(currentNode, "Z") {
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

	return steps
}

func gcf(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(steps ...int) int {
	result := steps[0] * steps[1] / gcf(steps[0], steps[1])

	for i := 2; i < len(steps); i++ {
		result = lcm(result, steps[i])
	}

	return result
}
