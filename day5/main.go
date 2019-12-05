package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func getVal(prog []int, idx int) int {
	return prog[idx]
}

// opcode -> (instruction, modes)
func parseOpcode(opcode int) (int, []int) {
	s := strconv.Itoa(opcode)
	if len(s) <= 2 {
		return opcode, []int{}
	}
	
	i := s[len(s)-2:]
	ms := s[:2]

	var modes []int
	for _, m := range ms {
		mode, _ := strconv.Atoi(string(m))
		modes = append(modes, mode)
	}

	instr, _ := strconv.Atoi(i)
	return instr, modes
}

func get(prog []int, val, mode int) int {
	if mode == 1 {
		fmt.Printf("Get %d mode %d = %d\n", val, mode, prog[val])
		return prog[val]
	}

	fmt.Printf("Get %d mode %d = %d\n", val, mode, prog[prog[val]])
	return prog[prog[val]]
}

func getMode(modes []int, index int) int {
	if index < len(modes) {
		return modes[len(modes)-index-1]
	}

	return 0
}

func runProgram(numbers []int) []int {
	for i := 0; i < len(numbers); {
		//opCode := numbers[i]
		opCode, modes := parseOpcode(numbers[i])
		fmt.Printf("Opcode: %d Index %d Modes %v\n", opCode, i, modes)
		switch opCode {
		case 1:
			first := get(numbers, i+1, getMode(modes, 0))
			second := get(numbers, i+2, getMode(modes, 1))
			resultIdx := numbers[i + 3]
			sum := first + second

			numbers[resultIdx] = sum
			i += 4
		case 2:
			first := get(numbers, i+1, getMode(modes, 0))
			second := get(numbers, i+2, getMode(modes, 1))
			resultIdx := numbers[i + 3]
			prod := first * second

			fmt.Printf("Result index: %d First: %d Second: %d Prod: %d\n",
				resultIdx, first, second, prod)
			numbers[resultIdx] = prod
			i += 4
		case 3:
			firstIdx := numbers[i + 1]
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			nS := strings.TrimSuffix(string(text), "\n")
			n, _ := strconv.Atoi(nS)
			numbers[firstIdx] = n
			i += 2

		case 4:
			firstIdx := numbers[i + 1]
			val := numbers[firstIdx]
			fmt.Printf("Instruction Output: %d\n", val)
			i += 2
			
		case 99:
			break
		default:
			i++
			break
		}
	}
	return numbers
}

func main() {
	//	lol, modes := parseOpcode(1002)
	//fmt.Printf("inst %d modes %d\n", lol, modes)

	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")
	list := strings.FieldsFunc(niceData, func(c rune) bool {
		return c == ','
	})

	var numbers []int
	for _, elem := range list {
		n, _ := strconv.Atoi(elem)
		numbers = append(numbers, n)
	}
	
	output := runProgram(numbers)
	
	fmt.Printf("Output: %v\n", output)
}
