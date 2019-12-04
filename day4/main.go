package main

import (
	"fmt"
)

const start = 136818
const end = 685979

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	
	return sequence
}

func valid(digits []int) bool {
	if increase(digits) {
		if same(digits) {
			return true
		}
	}

	return false
	//return increase(digits) && same(digits)
}

func same(digits []int) bool {
	count := 0
	list := make(map[int]int)

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			//fmt.Printf("! true\n")
			list[digits[i]]++
		} else {
			//fmt.Printf("! false\n")
			//list[digits[i]] = false
		}
	}

	for k, v := range list {
		fmt.Printf("Set: %d: %v\n", k, v)
		if v == 1 {
			count++
		}
	}

	return count > 0
}

func increase(digits []int) bool {
	prev := digits[0]
	for i := 1; i < len(digits); i++ {
		if digits[i] < prev {
			return false
		}

		prev = digits[i]
	}

	return true
}

func main() {
	count := 0
	//lol := valid([]int{1,1,3,7,7,9})
	
	//fmt.Printf("%t\n", lol)
	
	for i := start; i <= end; i++ {
		var seq []int
		result := IntToSlice(i, seq)
		ok := valid(result)

		if ok {
			count++
			fmt.Printf("%d valid: %T\n", i, ok)
		}
	}

	fmt.Printf("Count %d\n", count)
}
