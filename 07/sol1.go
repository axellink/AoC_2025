package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

func Run(start_pos Pos, splitters map[Pos]bool, map_end int, cache map[Pos]struct{}) map[Pos]bool {
	_, cached := cache[start_pos]
	if cached {
		return splitters
	} else {
		cache[start_pos] = struct{}{}
	}

	curr_pos := start_pos
	_, err := splitters[curr_pos]
	for !err && curr_pos.y <= map_end {
		curr_pos.y++
		_, err = splitters[curr_pos]
	}
	if !err {
		return splitters
	} else {
		splitters[curr_pos] = true
		splitters = Run(Pos{curr_pos.x - 1, curr_pos.y}, splitters, map_end, cache)
		splitters = Run(Pos{curr_pos.x + 1, curr_pos.y}, splitters, map_end, cache)
		return splitters
	}
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
	splitters := make(map[Pos]bool, 0)
	var start_pos Pos
	y := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, n := range line {
			if n == 'S' {
				start_pos = Pos{x: i, y: y}
			}
			if n == '^' {
				splitters[Pos{x: i, y: y}] = false
			}
		}
		y++
	}

	Run(start_pos, splitters, y, make(map[Pos]struct{}))

	for _, n := range splitters {
		if n {
			count++
		}
	}

	fmt.Println(count)

}
