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

	ranges := make([][]int, 0)
	ids := make([]int, 0)
	in_ranges := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			in_ranges = false
		} else if in_ranges {
			r := make([]int, 2)
			s := strings.Split(data, "-")
			r[0], _ = strconv.Atoi(s[0])
			r[1], _ = strconv.Atoi(s[1])
			ranges = append(ranges, r)
		} else {
			id, _ := strconv.Atoi(data)
			ids = append(ids, id)
		}
	}

	for _, n := range ids {
		for _, r := range ranges {
			if n >= r[0] && n <= r[1] {
				count++
				break
			}
		}
	}

	fmt.Println(count)

}
