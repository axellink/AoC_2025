package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

//func max(s []rune) (int, int) {
//	index := 0
//	max := s[0]
//	for i, x := range s[1:] {
//		if x > max {
//			index = i + 1
//			max = x
//		}
//	}
//	return index, int(max - '0')
//}

func Sol2() {
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
		index := 0
		res := 0
		for i := 11; i >= 0; i-- {
			new_index, m := max(line[index : len(line)-i])
			res += m * int(math.Pow10(i))
			index += new_index + 1
		}
		count += res
	}

	fmt.Println(count)

}
