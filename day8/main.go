package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	_ "strconv"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	fmt.Printf("Nicedata: %s\n", niceData)
}
