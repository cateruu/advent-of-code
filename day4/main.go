package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Game struct {
	Won           int
	NumberOfCards int
}

func main() {
	start := time.Now()

	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)

	allCards := []Game{}
	for fs.Scan() {
		game := strings.Split(fs.Text(), ":")[1]

		numbers := strings.Split(game, "|")
		winningNumbers := strings.Split(numbers[0], " ")[1:]
		ourNumbers := strings.Split(numbers[1], " ")[1:]

		won := checkWinNumbers(ourNumbers, winningNumbers)
		allCards = append(allCards, Game{Won: won, NumberOfCards: 1})
	}

	for i, game := range allCards {
		for j := 0; j < game.NumberOfCards; j++ {
			playCopiedCards(&game, &allCards, game.Won, i)
		}
	}

	sum := 0
	for _, g := range allCards {
		sum += g.NumberOfCards
	}
	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
}

func checkWinNumbers(nums []string, winNums []string) int {
	won := 0
	for _, n := range nums {
		if slices.Contains(winNums, n) && n != "" {
			won++
		}
	}

	return won
}

func playCopiedCards(currentCard *Game, allCards *[]Game, cardsLeft int, currentCardIdx int) {
	if cardsLeft <= 0 || currentCardIdx+1 >= len(*allCards) {
		return
	}

	(*allCards)[currentCardIdx+1].NumberOfCards++

	playCopiedCards(currentCard, allCards, cardsLeft-1, currentCardIdx+1)
}
