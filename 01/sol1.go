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
	pos := 50
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation, _ := strconv.Atoi(strings.Replace(strings.Replace(scanner.Text(), "L", "-", 1), "R", "+", 1))
		pos = (rotation + pos) % 100
		if pos == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(count)
	}
}
