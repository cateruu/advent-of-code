package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Position struct {
	Line      int
	Index     int
	Direction string
	Pipe      rune
}

func main() {
	start := time.Now()
	readFile, err := os.Open("input.txt")
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
	fullLoopNumber := findFurthestPipe(startingPos, &pipeMap)

	elapsed := time.Since(start)
	fmt.Println(fullLoopNumber/2, elapsed)
}

func findFurthestPipe(startingPos Position, pipeMap *[]string) int {
	nextPipe := findNextPipeFromStart(startingPos, *pipeMap)

	fullLoopNumber := getFullLoopNumber(nextPipe, *pipeMap, startingPos)

	return fullLoopNumber
}

func getFullLoopNumber(pipePos Position, pipeMap []string, startPos Position) int {
	pipe := pipePos
	steps := 1
	for {
		pipe = findNextPipe(pipe, pipeMap)
		if pipe.Line == startPos.Line && pipe.Index == startPos.Index {
			break
		}

		steps++
	}

	// pipe = reversePipeDirection(pipePos)
	// i := 0
	// for {
	// 	// fmt.Println("curr", string(pipe.Pipe), pipe.Direction)
	// 	pipe = findNextPipe(pipe, pipeMap)
	// 	// fmt.Println("next", string(pipe.Pipe), pipe.Direction)

	// 	if pipe.Direction == "" {
	// 		break
	// 	}

	// 	if pipe.Line == startPos.Line && pipe.Index == startPos.Index {
	// 		break
	// 	}

	// 	i++
	// 	steps++
	// }
	// rank[1] = steps

	return steps + 1
}

// func reversePipeDirection(pipe Position) Position {
// 	direction := pipe.Direction
// 	switch pipe.Pipe {
// 	case '|':
// 		if direction == "n" {
// 			direction = "s"
// 		} else {
// 			direction = "n"
// 		}
// 	case '-':
// 		if direction == "w" {
// 			direction = "e"
// 		} else {
// 			direction = "w"
// 		}
// 	case 'L':
// 		if direction == "s" {
// 			direction = "w"
// 		} else {
// 			direction = "s"
// 		}
// 	case 'J':
// 		if direction == "s" {
// 			direction = "w"
// 		} else {
// 			direction = "s"
// 		}
// 	case '7':
// 		if direction == "n" {
// 			direction = "e"
// 		} else {
// 			direction = "n"
// 		}
// 	case 'F':
// 		if direction == "n" {
// 			direction = "w"
// 		} else {
// 			direction = "n"
// 		}
// 	}

// 	return Position{Line: pipe.Line, Index: pipe.Index, Direction: direction, Pipe: pipe.Pipe}
// }

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
			line := startingPos.Line - 1
			index := startingPos.Index
			pipe := rune(pipeMap[line][index])

			if pipe != '.' {
				return Position{Line: line, Index: index, Direction: "n", Pipe: pipe}
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
