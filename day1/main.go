package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt");
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)

	sum := 0
	for fs.Scan() {
		numbers := []string{} 

		for i := 0; i < len(fs.Text()); i++ {
			if fs.Text()[i] >= '0' && fs.Text()[i] <= '9' {
				numbers = append(numbers, string(fs.Text()[i]))
			}
		}

		first := numbers[0]
		last := numbers[len(numbers) - 1]
		
		number, _ := strconv.Atoi(first + last)

		sum += number
	}

	fmt.Println(sum)
}