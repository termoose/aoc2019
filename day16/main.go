package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func BasePattern() []int {
	return []int{0, 1, 0, -1}
}

func getPattern(n, size int) []int {
	var pattern []int
	base := BasePattern()

	for _, b := range base {
		for i := 0; i <= n; i++ {
			pattern = append(pattern, b)
		}
	}

	var result []int
	for i := 0; i <= size; i++ {
		result = append(result, pattern[i % len(pattern)])
	}

	return result
}

func applyPattern(n int, numbers []int) int {
	pattern := getPattern(n, len(numbers))[1:]
	result := 0

	for i, n := range numbers {
		p := pattern[i % len(pattern)]
		result += (n * p) % 10
	}

	return int(math.Abs(float64(result % 10)))
}

func parse() []int {
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")
	var result []int

	for _, d := range niceData {
		n, _ := strconv.Atoi(string(d))
		result = append(result, n)
	}

	return result
}

func iterate(list []int, n int) []int {
	if n == 0 {
		return list
	}

	var result []int

	for i := 0; i < len(list); i++ {
		n := applyPattern(i, list)
		result = append(result, n)
	}

	return iterate(result, n - 1)
}

func main() {
	list := parse()
	fmt.Printf("%v\n", list)
	
	result := iterate(list, 100)
	fmt.Printf("Result: %v\n", result[:8])
}
