package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SplitByN(s string, n int) []string {
	elements := len(s) / n
	if len(s)%n != 0 {
		elements++
	}
	res := make([]string, elements)
	offset := 0
	for i := 0; offset < len(s); i++ {
		end := offset + n
		if end >= len(s) {
			res[i] = string([]rune(s)[offset:])
		} else {
			res[i] = string([]rune(s)[offset:end])
		}
		offset = end
	}
	return res
}

func CheckReappearance(seq []string) bool {
	ln := len(seq)
	for i := 0; i < ln-1; i++ {
		el := seq[i]
		for _, x := range seq[i+1:] {
			if el != x {
				return false
			}
		}
		return true
	}
	return false
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
	scanner.Scan()
	data := scanner.Text()
	for _, s := range strings.Split(data, ",") {
		bounds := strings.Split(s, "-")
		lower, _ := strconv.Atoi(bounds[0])
		higher, _ := strconv.Atoi(bounds[1])
		for i := lower; i <= higher; i++ {
			str := strconv.Itoa(i)
			mid := len(str) / 2
			for j := 1; j <= mid; j++ {
				if CheckReappearance(SplitByN(str, j)) {
					count += i
					break
				}
			}
		}
	}

	fmt.Println(count)

}
