package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	"golang.org/x/exp/slices"
)

type Position struct {
	Line        int
	Index       int
	Direction   string
	Pipe        rune
	IsEdge      bool
	IsEndOfLine bool
}

func main() {
	start := time.Now()
	readFile, err := os.Open("test.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	pipeMap := []string{}
	for fs.Scan() {
		line := fs.Text()

		pipeMap = append(pipeMap, line)
	}

	startingPos := findStartingPos(pipeMap)
	allPipes := getAllPipesPositions(startingPos, &pipeMap)
	sort.Slice(allPipes, func(i, j int) bool {
		if allPipes[i].Line != allPipes[j].Line {
			return allPipes[i].Line < allPipes[j].Line
		}

		return allPipes[i].Index < allPipes[j].Index
	})
	allPipes = setUpEdges(allPipes)
	trapped := trappedByLoop(allPipes, pipeMap)

	elapsed := time.Since(start)
	fmt.Println(trapped, elapsed)
}

func trappedByLoop(loop []Position, pipeMap []string) int {
	trappedAmount := 0
	isTrapped := false
	for lineIdx, line := range pipeMap {
		for itemIdx := range line {
			fmt.Println(isTrapped)
			pipeIdx := slices.IndexFunc(loop, func(i Position) bool {
				return i.Line == lineIdx && i.Index == itemIdx
			})

			if pipeIdx != -1 {
				if loop[pipeIdx].IsEdge {
					isTrapped = !isTrapped
				}

				if loop[pipeIdx].IsEndOfLine {
					isTrapped = false
				}

				continue
			}

			if isTrapped {
				fmt.Println("here")
				trappedAmount++
			}
		}
		fmt.Println()
	}

	return trappedAmount
}

func setUpEdges(loop []Position) []Position {
	positions := make([]Position, len(loop))
	currLine := loop[0].Line
	wasOnLine := false
	for i, pos := range loop {
		if currLine != pos.Line {
			currLine = pos.Line
			wasOnLine = false
		}

		// bottom right pipe
		if i == len(loop)-1 {
			pos.IsEdge = true
		}

		// chcek right
		if i != len(loop)-1 && pos.Line == loop[i+1].Line && pos.Index != loop[i+1].Index-1 {
			pos.IsEdge = true
		}

		// check left
		if i != 0 && pos.Line == loop[i-1].Line && pos.Index-1 != loop[i-1].Index {
			pos.IsEdge = true
		}

		// first of line
		if !wasOnLine {
			pos.IsEdge = true
			wasOnLine = true
		}

		//last of line
		if i != len(loop)-1 && pos.Line != loop[i+1].Line {
			pos.IsEdge = true
			pos.IsEndOfLine = true
		}

		positions[i] = pos
	}

	return positions
}

func getAllPipesPositions(startingPos Position, pipeMap *[]string) []Position {
	nextPipe := findNextPipeFromStart(startingPos, *pipeMap)
	pipes := []Position{startingPos}

	pipes = append(pipes, getPipesPositionsFromStart(nextPipe, *pipeMap, startingPos)...)
	return pipes
}

func getPipesPositionsFromStart(pipePos Position, pipeMap []string, startPos Position) []Position {
	pipes := []Position{}
	pipe := pipePos
	for {
		pipes = append(pipes, pipe)
		pipe = findNextPipe(pipe, pipeMap)
		if pipe.Line == startPos.Line && pipe.Index == startPos.Index {
			break
		}
	}

	return pipes
}

func findNextPipe(currentPos Position, pipeMap []string) Position {
	var line, index int
	var pipe rune
	var direction string
	if currentPos.Pipe == '|' {
		switch currentPos.Direction {
		case "n":
			line = currentPos.Line - 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = currentPos.Direction
			break
		case "s":
			line = currentPos.Line + 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = currentPos.Direction
			break
		}
	} else if currentPos.Pipe == '-' {
		switch currentPos.Direction {
		case "w":
			line = currentPos.Line
			index = currentPos.Index - 1
			pipe = rune(pipeMap[line][index])
			direction = currentPos.Direction
			break
		case "e":
			line = currentPos.Line
			index = currentPos.Index + 1
			pipe = rune(pipeMap[line][index])
			direction = currentPos.Direction
			break
		}
	} else if currentPos.Pipe == 'L' {
		switch currentPos.Direction {
		case "w":
			line = currentPos.Line - 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		case "s":
			line = currentPos.Line
			index = currentPos.Index + 1
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		}
	} else if currentPos.Pipe == 'J' {
		switch currentPos.Direction {
		case "e":
			line = currentPos.Line - 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		case "s":
			line = currentPos.Line
			index = currentPos.Index - 1
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		}
	} else if currentPos.Pipe == '7' {
		switch currentPos.Direction {
		case "e":
			line = currentPos.Line + 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		case "n":
			line = currentPos.Line
			index = currentPos.Index - 1
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		}
	} else if currentPos.Pipe == 'F' {
		switch currentPos.Direction {
		case "w":
			line = currentPos.Line + 1
			index = currentPos.Index
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		case "n":
			line = currentPos.Line
			index = currentPos.Index + 1
			pipe = rune(pipeMap[line][index])
			direction = getPipeDirection(currentPos.Pipe, currentPos.Direction)
			break
		}
	}

	return Position{Line: line, Index: index, Pipe: pipe, Direction: direction}
}

func getPipeDirection(pipe rune, prevDirection string) string {
	switch pipe {
	case '|':
		return prevDirection
	case '-':
		return prevDirection
	case 'L':
		if prevDirection == "w" {
			return "n"
		} else if prevDirection == "s" {
			return "e"
		}
	case 'J':
		if prevDirection == "e" {
			return "n"
		} else if prevDirection == "s" {
			return "w"
		}
	case '7':
		if prevDirection == "e" {
			return "s"
		} else if prevDirection == "n" {
			return "w"
		}
	case 'F':
		if prevDirection == "w" {
			return "s"
		} else if prevDirection == "n" {
			return "e"
		}
	}

	return ""
}

func findNextPipeFromStart(startingPos Position, pipeMap []string) Position {
	for i := 0; i < 4; i++ {
		if i == 0 {
			// line := startingPos.Line - 1
			// index := startingPos.Index
			line := startingPos.Line
			index := startingPos.Index + 1
			pipe := rune(pipeMap[line][index])

			if pipe != '.' {
				return Position{Line: line, Index: index, Direction: "e", Pipe: pipe}
			}
		} else if i == 1 {
			line := startingPos.Line
			index := startingPos.Index + 1
			pipe := rune(pipeMap[line][index])

			if pipe != '.' {
				return Position{Line: line, Index: index, Direction: "e", Pipe: pipe}
			}
		} else if i == 2 {
			line := startingPos.Line - 1
			index := startingPos.Index
			pipe := rune(pipeMap[line][index])

			if pipe != '.' {
				return Position{Line: line, Index: index, Direction: "s", Pipe: pipe}
			}
		} else if i == 3 {
			line := startingPos.Line + 1
			index := startingPos.Index - 1
			pipe := rune(pipeMap[line][index])

			if pipe != '.' {
				return Position{Line: line, Index: index, Direction: "w", Pipe: pipe}
			}
		}
	}

	return startingPos
}

func findStartingPos(pipeMap []string) Position {
	startingPos := Position{}
	for lineIdx, line := range pipeMap {
		for pipeIdx, pipe := range line {
			if pipe == 'S' {
				startingPos = Position{Line: lineIdx, Index: pipeIdx}
				break
			}
		}
	}

	return startingPos
}
