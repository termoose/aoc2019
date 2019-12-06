package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type pair struct {
	left  string
	right string
}

type node struct {
	planet   string
	children []*node
	depth    int
}

func getNode(root *node, planet string, depth int) (*node, int) {
	if root.planet == planet {
		return root, depth
	}

	for _, child := range root.children {
		result, d := getNode(child, planet, depth + 1)

		if result != nil {
			return result, d
		}
	}

	return nil, depth
}

func AddChild(root *node, parent, child string) {
	n, depth := getNode(root, parent, 0)

	if n != nil {
		new := &node{
			planet: child,
			children: []*node{},
			depth: depth + 1,
		}

		fmt.Printf("Adding %s under %s\n", child, parent)
		n.children = append(n.children, new)
	}
}

func orbitParse(line string) (string, string) {
	list := strings.FieldsFunc(line, func(c rune) bool {
		return c == ')'
	})

	return list[0], list[1]
}

func sumAll(root *node) int {
	if root == nil {
		return 0
	}

	fmt.Printf("Sum %s: %d Children: %V\n", root.planet, root.depth, root.children)

	sum := 0
	for _, child := range root.children {
		sum += sumAll(child)
	}

	return root.depth + sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := &node{
		planet: "COM",
		children: []*node{},
	}

	planets := []pair{}
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		first, second := orbitParse(line)

		planets = append(planets, pair{
			left: first,
			right: second,
		})

		//fmt.Printf("Adding %s around %s\n", second, first)
		//AddChild(root, first, second)
	}

	for planet := range planets {
	}

	node, _ := getNode(root, "COM", 0)
	fmt.Printf("Root: %v Child: %v Child: %v\n", node, node.children[0], node.children[0].children[0])
	sum := sumAll(root)
	fmt.Printf("Sum: %d\n", sum)
}
