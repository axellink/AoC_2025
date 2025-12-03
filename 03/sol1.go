package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func max(s []rune) (int, int) {
	index := 0
	max := s[0]
	for i, x := range s[1:] {
		if x > max {
			index = i+1
			max = x
		}
	}
	return index, int(max - '0')
}

func Sol1() {
	filename := os.Args[1]
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		i, m := max(line[:len(line)-1])
		_, m2 := max(line[i+1:])
		count += m*10 + m2
	}

	fmt.Println(count)

}
