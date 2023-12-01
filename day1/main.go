package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		numbers := map[int]string{}
		indexes := []int{}

		str := fs.Text()
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				numbers[i] = string(str[i])
				indexes = append(indexes, i)
			}
		}

		nums := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
		for k, v := range nums {
			if strings.Contains(str, k) {
				idx := strings.Index(str, k)
				indexes = append(indexes, idx)
				numbers[idx] = v

				if strings.Count(str, k) > 1 {
					lastIdx := strings.LastIndex(str, k)
					indexes = append(indexes, lastIdx)
					numbers[lastIdx] = v
				}
			}
		}

		sort.Ints(indexes)

		first := numbers[indexes[0]]
		last := numbers[indexes[len(indexes)-1]]

		number, _ := strconv.Atoi(first + last)

		sum += number
	}

	fmt.Println(sum)
}
