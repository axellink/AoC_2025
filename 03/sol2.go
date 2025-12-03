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

// not sure that doing this conversion is faster, equal or slower than
// index calculation in previous commit but I wanted to try to use
// slice len and cap to my advantage in this logic so I needed a well
// sized slice, which scanner.Text() does not do because of buffering
// I got a rune slice with larger cap, messing up my logic
func convert_to_rune_with_cap(s string) []rune {
	res := make([]rune, len(s))
	copy(res, []rune(s))
	return res
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
	for scanner.Scan() {
		line := convert_to_rune_with_cap(scanner.Text())
		index := 0
		m := 0
		for i := 11; i >= 0; i-- {
			line = line[index : cap(line)-i]
			index, m = max(line)
			count += m * int(math.Pow10(i))
			index++
		}
	}

	fmt.Println(count)

}
