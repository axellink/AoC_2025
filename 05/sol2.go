package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(s string) []int {
	nums := strings.Split(s, "-")
	res := make([]int, 2)
	res[0], _ = strconv.Atoi(nums[0])
	res[1], _ = strconv.Atoi(nums[1])
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
	real_ranges := make([][]int, 0)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "" {
			break
		} else {
			curr_range := parse(data)
			var lower_placed_at int
			for lower_placed_at = 0; lower_placed_at < len(real_ranges); lower_placed_at++ {
				if curr_range[0] <= real_ranges[lower_placed_at][1] {
					break
				}
			}
			if lower_placed_at >= len(real_ranges) { // check has to be appended
				real_ranges = append(real_ranges, curr_range)
			} else if curr_range[0] >= real_ranges[lower_placed_at][0] && curr_range[1] <= real_ranges[lower_placed_at][1] { //check completely contained
				continue
			} else if curr_range[1] < real_ranges[lower_placed_at][0] { // check completely before
				real_ranges = slices.Insert(real_ranges, lower_placed_at, curr_range)
			} else { // it's not contained nor before, one or both bound is outside
				if curr_range[0] < real_ranges[lower_placed_at][0] { // if it's the lower bound, just change it
					real_ranges[lower_placed_at][0] = curr_range[0]
				}
				if curr_range[1] > real_ranges[lower_placed_at][1]{ // it's the upper bound, so we have to change it, and maybe update the other bounds
					// first let's find where our upper bound lays (before or inside an other bounds)
					rest := real_ranges[lower_placed_at:]
					for len(rest) > 0 && curr_range[1] > rest[0][1] {
						rest = rest[1:]
					}
					if len(rest) == 0 { // we were higher than all other bounds, update upper bound an get rid of the old rest of the ranges
						real_ranges[lower_placed_at][1] = curr_range[1]
						real_ranges = real_ranges[:lower_placed_at+1]
					} else if curr_range[1] < rest[0][0] { // our new upper bound is before the next range, so we need to get rid of ranges between, and update upper bound
						real_ranges[lower_placed_at][1] = curr_range[1]
						real_ranges = append(real_ranges[:lower_placed_at+1], rest...)
					} else { // our new upper bound is insied the next range, so the real new upper bound is this range upper bound, and we need to get rid of ranges in between
						real_ranges[lower_placed_at][1] = rest[0][1]
						real_ranges = append(real_ranges[:lower_placed_at+1], rest[1:]...)
					}
				}
			}
		}
	}

	for _, r := range real_ranges {
		count += r[1] - r[0] + 1
	}

	fmt.Println(count)

}
