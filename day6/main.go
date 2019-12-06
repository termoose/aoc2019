package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type node struct {
	planet   string
	children []*node
	depth    int
}


func printNode(root *node) {
	fmt.Printf("Node %s Children: ", root.planet)
	for _, c := range root.children {
		fmt.Printf("%s ", c.planet)
	}
	fmt.Println()
}

func printTree(root *node) {
	if root == nil {
		return
	}

	for _, c := range root.children {
		printNode(c)
		printTree(c)
	}
}

func getDepth(root *node, planet string) int {
	if root == nil {
		return -1
	}

	for _, c := range root.children {
		fmt.Printf("Checking %s == %s\n", c.planet, planet)
		//printNode(c)
		
		if c.planet == planet {
			return c.depth
		}
		
		return getDepth(c, planet)
	}

	return root.depth
}

func getNode(root *node, planet string, depth int) (*node, int) {
	if root == nil {
		return nil, depth
	}

	if root.planet == planet {
		return root, depth
	}

	for _, child := range root.children {
		return getNode(child, planet, depth + 1)
	}

	return root, depth
}

func AddChild(root *node, parent, child string) {
	n, depth := getNode(root, parent, 0)

	if n != nil {
		new := &node{
			planet: child,
			children: []*node{},
			depth: depth + 1,
		}

		fmt.Printf("Adding %s to %s\n", child, n.planet)
		n.children = append(n.children, new)
	} else {
		fmt.Printf("Could not find parent %s\n", parent)
	}
}

func orbitParse(line string) (string, string) {
	list := strings.FieldsFunc(line, func(c rune) bool {
		return c == ')'
	})

	return list[0], list[1]
}

func output(node *node) {
	if node == nil {
		return
	}

	for _, child := range node.children {
		output(child)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := &node{
		planet: "COM",
		children: []*node{},
	}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		first, second := orbitParse(line)

		fmt.Printf("Adding %s around %s\n", second, first)
		AddChild(root, first, second)

		// node, _ := getNode(root, first, 0)
		// if node != nil {
			// c := count(node)
			// fmt.Printf("Count %s: %d\n", node.planet, c)
			// fmt.Printf("Depth %d %s children: ", depth, node.planet)
			// for _, c := range node.children {
			// 	fmt.Printf("%s ", c.planet)
			// }
			// fmt.Println()
		// }
	}

	fmt.Printf("-----\n")
	printTree(root)
	fmt.Printf("-----\n")
	depth := getDepth(root, "G")
	fmt.Printf("depth %d\n", depth)
}
