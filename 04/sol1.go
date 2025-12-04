package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func count_around(maze [][]rune, x int, y int) int {
	dirs := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	count := 0
	for _, d := range dirs {
		_x := x + d[0]
		_y := y + d[1]
		if _y >= 0 && _x >= 0 && _y < len(maze) && _x < len(maze[y]) && maze[_y][_x] == '@' {
			count++
		}
	}
	return count
}

func Sol1() {
	filename := os.Args[1]
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maze := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, []rune(scanner.Text()))
	}

	for y, r := range maze {
		for x, c := range r {
			if c == '@' {
				if count_around(maze, x, y) < 4 {
					count++
				}
			}
		}
	}

	fmt.Println(count)

}
