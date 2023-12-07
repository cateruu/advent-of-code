package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/icza/abcsort"
)

type Hand struct {
	Hand string
	Bid  int
}

func main() {
	start := time.Now()
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		panic(err.Error())
	}

	fs := bufio.NewScanner(readFile)

	handTypes := make(map[string][]Hand)
	for fs.Scan() {
		game := strings.Split(fs.Text(), " ")

		handType := getHandType(game[0])
		bid, _ := strconv.Atoi(game[1])
		handTypes[handType] = append(handTypes[handType], Hand{Hand: game[0], Bid: bid})
	}

	sorter := abcsort.New("J123456789TQKA")
	for _, hand := range handTypes {
		sorter.Slice(hand, func(i int) string {
			return hand[i].Hand
		})
	}

	typeSlice := []string{"HighCard", "OnePair", "TwoPair", "ThreeOfKind", "FullHouse", "FourOfKind", "FiveOfKind"}

	sum := 0
	rank := 1
	for _, handType := range typeSlice {
		for _, hand := range handTypes[handType] {
			sum += hand.Bid * rank
			rank++
		}
	}

	elapsed := time.Since(start)
	fmt.Println(sum, elapsed)
}

func getHandType(hand string) string {
	allCards := make(map[rune]int)
	for _, card := range hand {
		allCards[card]++
	}

	var biggestPair rune
	biggestValue := 0
	jokerValue := 0
	for k, v := range allCards {
		if k == 'J' {
			jokerValue = v
			continue
		}

		if v > biggestValue {
			biggestPair = k
			biggestValue = v
		}
	}

	_, ok := allCards['J']
	if ok {
		allCards[biggestPair] += jokerValue
		delete(allCards, 'J')
	}

	nums := []string{}
	for _, card := range allCards {
		nums = append(nums, fmt.Sprint(card))
	}

	sort.Strings(nums)
	handValue := strings.Join(nums, "")

	if handValue == "5" {
		return "FiveOfKind"
	} else if handValue == "14" {
		return "FourOfKind"
	} else if handValue == "23" {
		return "FullHouse"
	} else if handValue == "113" {
		return "ThreeOfKind"
	} else if handValue == "122" {
		return "TwoPair"
	} else if handValue == "1112" {
		return "OnePair"
	} else {
		return "HighCard"
	}
}
