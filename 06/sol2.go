package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"strings"
)

func Sol2() {
	filename := os.Args[1]
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([][]rune, 0)
	results := make([]int, 0)
	lines := 0
	for scanner.Scan() {
		data := []rune(scanner.Text())
		if data[0] != '+' && data[0] != '*' {
			lines++
			numbers = append(numbers, data)
		} else {
			for true {
				op := data[0]
				num_space := 1
				for len(data) > num_space && data[num_space] == ' ' {
					num_space++
				}
				end := false
				if len(data) == num_space {
					num_space++
					end = true
				}
				tmp_numbers := make([][]rune, num_space-1)
				for i := range tmp_numbers {
					tmp_numbers[i] = make([]rune, lines)
					for j := range tmp_numbers[i] {
						tmp_numbers[i][j] = numbers[j][0]
						numbers[j] = numbers[j][1:]
					}
				}
				switch op {
				case '+':
					res := 0
					for _, r := range tmp_numbers {
						n, _ := strconv.Atoi(strings.TrimSpace(string(r)))
						res += n
					}
					results = append(results, res)
				case '*':
					res := 1
					for _, r := range tmp_numbers {
						n, _ := strconv.Atoi(strings.TrimSpace(string(r)))
						res *= n
					}
					results = append(results, res)
				}
				if !end {
					data = data[num_space:]
					for i := range numbers { // delete space left on numbers lines
						numbers[i] = numbers[i][1:]
					}
				} else {
					break
				}
			}
		}
	}

	for _, n := range results {
		count += n
	}

	fmt.Println(count)

}
