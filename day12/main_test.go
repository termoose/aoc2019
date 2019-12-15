package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestFirstPart(t *testing.T) {
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	moons := parse(niceData)

	for i := 0; i < 1000; i++ {
		iterate(&moons)
	}

	energy := moons.energy()
	if energy != 6735 {
		t.Errorf("Energy %d not equal to 6735", energy)
	}
}