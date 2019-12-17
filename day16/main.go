package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func BasePattern() []int {
	return []int{ 0, 1, 0, -1 }
}

func calcPatternVal(n, i, offset int) int {
	base := BasePattern()

	return base[((i+offset) / (n+1)) % len(base)]
}

func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func applyPattern(size int, numbers []int) int {
	result := 0

	// Diagonal
	for i := 0; i < len(numbers); i++ {
		p := calcPatternVal(size, i, 1)
		n := numbers[i]
		fmt.Printf("%d*%d + ", n, p)
		result += (n * p) % 10
	}
	fmt.Printf("= %d\n", abs(result % 10))

	return abs(result % 10)
}

func applyBigPattern(size int, numbers []int) int {
	result := 0

	for i := size; i < len(numbers); i++ {
		p := calcPatternVal(size, i, 1)
		n := numbers[i]
		fmt.Printf("%d*%d + ", n, p)
		result += (n * p) % 10
	}
	fmt.Printf("= %d DIAGONAL \n", abs(result % 10))

	upperPart := 0
	for i := 0; i < len(numbers); i++ {
		p := calcPatternVal(size, i, 1)
		n := numbers[i]
		fmt.Printf("%d*%d + ", n, p)
		upperPart += (n * p) % 10
	}

	fmt.Printf("= %d TOP\n", abs(result % 10))

	return abs((result + upperPart) % 10)
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

	fmt.Printf("Iteration: %d\n", n)

	return iterate(result, n - 1)
}

func iterateBig(list []int, n int) []int {
	if n == 0 {
		return list
	}

	// Diagonal
	var result []int
	for i := 0; i < len(list); i++ {
		line := applyBigPattern(i, list)
		result = append(result, line)
	}

	return iterateBig(result, n - 1)
}

func list2int(list []int) int {
	result := 0

	for i, n := range list {
		result += int(math.Pow10(len(list) - 1 - i)) * n
	}

	return result
}

func parse() []int {
	data, _ := ioutil.ReadFile("lol.txt")
	niceData := strings.TrimSuffix(string(data), "\n")
	var result []int

	for _, d := range niceData {
		n, _ := strconv.Atoi(string(d))
		result = append(result, n)
	}

	return result
}

func parseBig() []int {
	var result []int
	smol := parse()

	for i := 0; i < 2; i++ {
		result = append(result, smol...)
	}

	return result
}

func main() {
	list := parse()
	fmt.Printf("%v\n", list)

	big := parseBig()
	result := iterate(big, 1)
	fmt.Printf("Real result: %d\n", list2int(result[:8]))

	fast := iterateBig(list, 1)
	fmt.Printf("Result?: %d\n", list2int(fast[:8]))
	//offset := list2int(list[:7])
	//fmt.Printf("Offset: %d\n", offset)

	//big := parseBig()
	//fmt.Printf("Big: %d Small: %d\n", len(big), len(list))
	//result := iterate(big, 100)
	//message := result[offset:offset+8]
	////result := iterate(list, 100)
	//fmt.Printf("Message: %v\n", message)
}
