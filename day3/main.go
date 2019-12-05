package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strings"
	"strconv"
)

type intersection struct {
	x     int
	y     int
	steps int
}

type wire struct {
	visited int
	length  int
}

const size int = 25500

func move(actions []string) [][]wire {
	grid := make([][]wire, size)
	for i, _ := range grid {
		grid[i] = make([]wire, size)
	}
	currX := size / 2
	currY := size / 2
	counter := 0
	for _, move := range actions {
		switch string(move[0]) {
		case "L":
			length, _ := strconv.Atoi(move[1:])
			for i := 0; i < length; i++ {
				currX--
				counter++
				//l := grid[currY][currX].length

				grid[currY][currX].visited = 1
				grid[currY][currX].length = counter
			}

		case "R":
			length, _ := strconv.Atoi(move[1:])
			for i := 0; i < length; i++ {
				currX++
				counter++
				//l := grid[currY][currX].length

				grid[currY][currX].visited = 1
				grid[currY][currX].length = counter
			}

		case "U":
			length, _ := strconv.Atoi(move[1:])
			for i := 0; i < length; i++ {
				currY--
				counter++
				//l := grid[currY][currX].length

				grid[currY][currX].visited = 1
				grid[currY][currX].length = counter
			}

		case "D":
			length, _ := strconv.Atoi(move[1:])
			for i := 0; i < length; i++ {
				currY++
				counter++
				//l := grid[currY][currX].length
				
				grid[currY][currX].visited = 1
				grid[currY][currX].length = counter
			}
		}

	}

	return grid
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		lines = append(lines, line)
	}

	grid := make([][][]wire, 2)
	sum := make([][]int, size)
	length := make([][]int, size)
	for i, _ := range sum {
		sum[i] = make([]int, size)
		length[i] = make([]int, size)
	}

	inter := []intersection{}

	for i, line := range lines {
		l := strings.FieldsFunc(line, func(c rune) bool {
		 	return c == ','
		})

		grid[i] = move(l)

		for i, col := range grid[i] {
			for j, row := range col {
				sum[i][j] += row.visited
				length[i][j] += row.length

				if sum[i][j] > 1 {
					inter = append(inter,
						intersection{
							x: i-size/2,
							y: j-size/2,
							steps: length[i][j],
						})
				}
				//fmt.Printf("(%2d,%2d) %v ", i,j,row)
				//fmt.Printf("%v", row)
			}
			//fmt.Println()
		}
	}

	smallest := intersection{steps: 100000000}
	for _, i := range inter {
		dist := math.Abs(float64(i.x)) + math.Abs(float64(i.y))
		if i.steps < smallest.steps {
			smallest = i
		}
		fmt.Printf("Intersection %v dist: %f\n", i, dist)
	}

	fmt.Printf("Smallest: %v\n", smallest)
	
	// for _, col := range sum {
	// 	for _, row := range col {
	// 		fmt.Printf("%v", row)
	// 	}
	// 	fmt.Println()
	// }
}
