package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	left  string
	right string
}

type children struct {
	child map[string]bool
}

type universe struct {
	planets map[string]children
}

func orbitParse(line string) pair {
	list := strings.FieldsFunc(line, func(c rune) bool {
		return c == ')'
	})

	return pair{list[0], list[1]}
}

func path(tree universe, current, node string, trail []string) []string {
	trail = append(trail, current)

	if current == node {
		return trail
	}

	currentNode := tree.planets[current]
	result := []string{}
	for c, _ := range currentNode.child {
		result = append(result, path(tree, c, node, trail)...)
	}

	return result
}

func lca(tree universe, first, second string) string {
	p1 := path(tree, "COM", first, []string{})
	p2 := path(tree, "COM", second, []string{})

	for i, p := range p1 {
		if i < len(p2) {
			if p != p2[i] {
				return p2[i-1]
			}
		}
	}

	if len(p1) > len(p2) {
		return p2[len(p2)-1]
	}

	return p1[len(p1)-1]
}

func dist(tree universe, first, second string) int {
	distFirst := path(tree, "COM", first, []string{})
	distSecond := path(tree, "COM", second, []string{})
	LCA := lca(tree, first, second)
	distLCA := path(tree, "COM", LCA, []string{})

	return len(distFirst) + len(distSecond) - 2 * len(distLCA)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	planets := []pair{}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		planet := orbitParse(line)

		planets = append(planets, planet)
	}

	var tree universe
	tree.planets = make(map[string]children)
	for _, p := range planets {
		_, exists := tree.planets[p.left]

		if exists {
			tree.planets[p.left].child[p.right] = true
		} else {
			tree.planets[p.left] = children{
				child: make(map[string]bool),
			}
			tree.planets[p.left].child[p.right] = true
		}
	}

	d := dist(tree, "YOU", "SAN")
	fmt.Printf("Dist: %d\n", d - 2)

}
