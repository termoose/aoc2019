package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func runProgram(numbers []int, noun, verb int) int {
	numbers[1] = noun
	numbers[2] = verb

	for i := 0; i < len(numbers); i += 4 {
		switch numbers[i] {
		case 1:
			firstIdx := numbers[i + 1]
			secondIdx := numbers[i + 2]
			resultIdx := numbers[i + 3]
			sum := numbers[firstIdx] + numbers[secondIdx]

			numbers[resultIdx] = sum
		case 2:
			firstIdx := numbers[i + 1]
			secondIdx := numbers[i + 2]
			resultIdx := numbers[i + 3]
			prod := numbers[firstIdx] * numbers[secondIdx]

			numbers[resultIdx] = prod
		case 99:
			break
		}
	}
	return numbers[0]
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")
	list := strings.FieldsFunc(niceData, func(c rune) bool {
		return c == ','
	})

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			var numbers []int
			for _, elem := range list {
				n, _ := strconv.Atoi(elem)
				numbers = append(numbers, n)
			}
			
			output := runProgram(numbers, noun, verb)

			if output == 19690720 {
				fmt.Printf("(%d, %d)\n", noun, verb)
				fmt.Printf("Result: %d\n", 100 * noun + verb)
			}
		}
	}

	//fmt.Printf("Output: %v\n", output)
}
