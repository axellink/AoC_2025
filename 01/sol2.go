package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Sol2() {
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
		non_corrected_pos := (rotation + pos)
		count += int(math.Abs(float64(non_corrected_pos / 100)))
		if non_corrected_pos <= 0 && pos != 0 {
			count += 1
		}
		pos = non_corrected_pos % 100
		if pos < 0 {
			pos += 100
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(count)
	}
}
