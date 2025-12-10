package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Sol1() {
	filename := os.Args[1]
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := strings.Fields(scanner.Text())
	numbers := make([][]int, len(data))
	results := make([]int, len(data))
	for i, n := range data {
		numbers[i] = make([]int, 0)
		num, _ := strconv.Atoi(n)
		numbers[i] = append(numbers[i], num)
	}
	for scanner.Scan() {
		data := strings.Fields(scanner.Text())
		for i, n := range data {
			num, err := strconv.Atoi(n)
			if err != nil {
				break
			}
			numbers[i] = append(numbers[i], num)
		}
		for i, op := range data {
			switch op {
			case "+":
				results[i] = 0
				for _, n := range numbers[i] {
					results[i] += n
				}
			case "*":
				results[i] = 1
				for _, n := range numbers[i] {
					results[i] *= n
				}
			}
		}
	}

	for _, n := range results {
		count += n
	}

	fmt.Println(count)

}
