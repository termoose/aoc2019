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

func hasChild(root *node, child string) bool {
	for _, c := range root.children {
		if c.planet == child {
			return true
		}
	}

	return false
}

func addChild(root *node, parent, child string, planets []pair) {
	n, depth := getNode(root, parent, 0)

	if n != nil {
		if hasChild(n, child) {
			return
		}

		new := &node{
			planet: child,
			children: []*node{},
			depth: depth + 1,
		}
		
		fmt.Printf("Adding %s under %s\n", child, parent)
		n.children = append(n.children, new)
	} else {
		// Find the parent and add it
		for _, planet := range planets {
			if parent == planet.right {
				fmt.Printf("Non-existing parent %s added\n", parent)
				addChild(root, planet.left, parent, planets)
			}
		}
	}
}

func orbitParse(line string) (string, string) {
	list := strings.FieldsFunc(line, func(c rune) bool {
		return c == ')'
	})

	return list[0], list[1]
}

func sum(root *node, total int) int {
	if root == nil {
		return 0
	}

	for _, c := range root.children {
		fmt.Printf("Node %s Total sum %d\n", c.planet, total)
		total = sum(c, total + 1)
		//return total
	}

	return total
}

func sumAll(root *node) int {
	if root == nil {
		return 0
	}

	//fmt.Printf("Sum %s: %d Children: %V\n", root.planet, root.depth, root.children)

	sum := 0
	for _, child := range root.children {
		sum += sumAll(child)
	}

	return root.depth + sum
}

func main() {
	file, _ := os.Open("lol.txt")
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
	}

	for _, planet := range planets {
		addChild(root, planet.left, planet.right, planets)
	}

	node, _ := getNode(root, "COM", 0)
	fmt.Printf("Root: %v Child: %v Child: %v\n", node, node.children[0], node.children[0].children[0])
	sum := sum(root, 0)
	fmt.Printf("Sum: %d\n", sum)
}
