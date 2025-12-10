package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//type Pos struct {
//	x int
//	y int
//}

func Run2(start_pos Pos, splitters map[Pos]int, map_end int) int {
	curr_pos := start_pos
	_, err := splitters[curr_pos]
	for !err && curr_pos.y <= map_end {
		curr_pos.y++
		_, err = splitters[curr_pos]
	}
	if !err {
		return 1
	} else {
		if splitters[curr_pos] != 0 {
			return splitters[curr_pos]
		}
		splitters[curr_pos] += Run2(Pos{curr_pos.x - 1, curr_pos.y}, splitters, map_end)
		splitters[curr_pos] += Run2(Pos{curr_pos.x + 1, curr_pos.y}, splitters, map_end)
		return splitters[curr_pos]
	}
}

func Sol2() {
	filename := os.Args[1]
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	splitters := make(map[Pos]int, 0)
	var start_pos Pos
	y := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, n := range line {
			if n == 'S' {
				start_pos = Pos{x: i, y: y}
			}
			if n == '^' {
				splitters[Pos{x: i, y: y}] = 0
			}
		}
		y++
	}

	count = Run2(start_pos, splitters, y)

	fmt.Println(count)

}
