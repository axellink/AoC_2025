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
	data := scanner.Text()
	for _, s := range strings.Split(data, ",") {
		bounds := strings.Split(s, "-")
		lower, _ := strconv.Atoi(bounds[0])
		higher, _ := strconv.Atoi(bounds[1])
		for i := lower; i <= higher; i++ {
			str := strconv.Itoa(i)
			mid := len(str) / 2
			first := string([]rune(str)[:mid])
			second := string([]rune(str)[mid:])
			if first == second {
				count += i
			}
		}
	}

	fmt.Println(count)

}
