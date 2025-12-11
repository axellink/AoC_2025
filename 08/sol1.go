package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
	z int
}

type Distance struct {
	distance float64
	p1       Pos
	p2       Pos
}

func CalcDistance(p1 Pos, p2 Pos) Distance {
	dist := math.Sqrt(
		math.Pow(float64(p1.x)-float64(p2.x), 2) +
			math.Pow(float64(p1.y)-float64(p2.y), 2) +
			math.Pow(float64(p1.z)-float64(p2.z), 2))
	return Distance{dist, p1, p2}
}

func CompDistance(d1 Distance, d2 Distance) int {
	return int(d1.distance) - int(d2.distance)
}

func ParsePos(s string) Pos {
	numbers := strings.Split(s, ",")
	x, _ := strconv.Atoi(numbers[0])
	y, _ := strconv.Atoi(numbers[1])
	z, _ := strconv.Atoi(numbers[2])
	return Pos{x, y, z}
}

type Circuit struct {
	circuit map[Pos]struct{}
}

func CreateCircuit(p1 Pos, p2 Pos) Circuit {
	c := make(map[Pos]struct{})
	c[p1] = struct{}{}
	c[p2] = struct{}{}
	return Circuit{c}
}

func (c *Circuit) Add(p Pos) {
	c.circuit[p] = struct{}{}
}

func (c *Circuit) IsIn(p Pos) bool {
	for i := range c.circuit {
		if i == p {
			return true
		}
	}
	return false
}

func (c *Circuit) Merge(c2 Circuit) {
	for i := range c2.circuit {
		c.circuit[i] = struct{}{}
	}
}

func CompCircuit(c1 Circuit, c2 Circuit) int {
	return len(c2.circuit) - len(c1.circuit) // negative for reverse sort
}

func Sol1() {
	filename := os.Args[1]
	iter, _ := strconv.Atoi(os.Args[2])
	count := 1

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	positions := make([]Pos, 0)
	for scanner.Scan() {
		positions = append(positions, ParsePos(scanner.Text()))
	}

	distances := make([]Distance, 0)
	for i, p := range positions[:len(positions)-1] {
		for _, p2 := range positions[i+1:] {
			distances = append(distances, CalcDistance(p, p2))
		}
	}

	slices.SortFunc(distances, CompDistance)
	distances = distances[:iter]
	circuits := make([]Circuit, 0)

	for _, n := range distances {
		i1 := -1
		i2 := -1
		for i, c := range circuits {
			if c.IsIn(n.p1) {
				i1 = i
			}
			if c.IsIn(n.p2) {
				i2 = i
			}
			if i1 != -1 && i2 != -1 {
				break
			}
		}
		if i1 == -1 && i2 == -1 {
			circuits = append(circuits, CreateCircuit(n.p1, n.p2))
		} else if i1 != -1 && i2 == -1 {
			circuits[i1].Add(n.p2)
		} else if i1 == -1 && i2 != -1 {
			circuits[i2].Add(n.p1)
		} else if i1 != -1 && i2 != -1 {
			if i1 != i2 {
				circuits[i1].Merge(circuits[i2])
				circuits = slices.Delete(circuits, i2, i2+1)
			}
		}
	}

	slices.SortFunc(circuits, CompCircuit)

	for _, c := range circuits[:3] {
		count *= len(c.circuit)
	}

	fmt.Println(count)

}
