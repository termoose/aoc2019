package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	_ "io/ioutil"
)

func calcFuel(value int) int {
	trans := int(value / 3) - 2

	if trans < 0 {
		return 0
	}

	return trans + calcFuel(trans)
}

func main() {
	lol := calcFuel(1969)
	fmt.Printf("lol: %d\n", lol)
	
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		//transformed := int(number / 3) - 2
		transformed := calcFuel(number)

		sum += transformed
	}

	fmt.Printf("Sum: %d\n", sum)
}
