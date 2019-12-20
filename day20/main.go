package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getNodeCoordinate(lines []string, node rune) (int, int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			line := lines[i]
			c := line[j]
			if rune(c) == node {
				if j - 1 >= 0 && line[j - 1] == '.' {
					return j - 1, i
				}
				if j + 1 < len(line) && line[j + 1] == '.' {
					return j + 1, i
				}
				if i - 1 >= 0 && lines[i - 1][j] == '.' {
					return j, i - 1
				}
				if i + 1 < len(lines) && lines[i + 1][j] == '.' {
					return j, i + 1
				}
			}
		}
	}

	return -1, -1
}

func getTeleport(lines []string, node rune) rune {
	nx, ny := getNodeCoordinate(lines, node)
	top := rune(lines[nx][ny - 1])
	left := rune(lines[nx - 1][ny])
	right := rune(lines[nx + 1][ny])
	bottom := rune(lines[nx][ny + 1])

	fmt.Printf("Top: %c Left: %c Right: %c Bottom: %c\n",
		top, left, right, bottom)

	if top != '.' && top != '#' && top != ' ' {
		return top
	}

	if left != '.' && left != '#' && left != ' ' {
		return left
	}

	if right != '.' && right != '#' && right != ' ' {
		return right
	}

	if bottom != '.' && bottom != '#' && bottom != ' ' {
		return bottom
	}

	return rune(0)
}

func getStart(lines []string) (int, int) {
	return getNodeCoordinate(lines, 'A')
}

func findReachables(lines []string, node rune) []rune {
	getNodeCoordinate(lines, node)

	return []rune{}
}

func main () {
	file, _ := os.Open("lol.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		lines = append(lines, line)
	}

	fmt.Printf("Lines: %d\n", len(lines))
	fmt.Printf("%v\n", lines)
	x, y := getStart(lines)
	fmt.Printf("Start %d, %d\n", x, y)

	cx, cy := getNodeCoordinate(lines, 'C')
	fmt.Printf("C coordinate: %d, %d\n", cx, cy)

	teleport := getTeleport(lines, 'C')
	fmt.Printf("C teleport to %c\n", teleport)
	//grid := make([][]rune, len(lines))
	//for i, _ := range grid {
	//	grid[i] = make([]rune, len(lines[0]) + 4)
	//}
	//
	//for y, line := range lines {
	//	for x, c := range line {
	//		grid[x][y] = c
	//	}
	//}
	//
	//fmt.Printf("%v\n", grid)
}
