package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type moon struct {
	x      int
	y      int
	z      int
	deltax int
	deltay int
	deltaz int
}

func (m *moon) setDeltas(other moon) {
	if m.x > other.x {
		m.deltax--
	} else {
		m.deltax++
	}

	if m.y > other.y {
		m.deltay--
	} else {
		m.deltay++
	}

	if m.z > other.z {
		m.deltaz--
	} else {
		m.deltaz++
	}
}

func parse(data string) []moon {
	var result []moon

	for _, line := range strings.Split(data, "\n") {
		var elem moon
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &elem.x, &elem.y, &elem.z)
		result = append(result, elem)
	}

	return result	
}

func applyGravity(moons []moon) []moon {
	var result []moon

	for _, m1 := range moons {
		for _, m2 := range moons {
			m1.setDeltas(m2)
		}

		result = append(result, m1)
	}

	return result
}

func printMoons(moons []moon) {
	for _, m := range moons {
		fmt.Printf("<x=%d, y=%d, z=%d> d: <%d, %d, %d>\n", m.x, m.y, m.z,
		m.deltax, m.deltay, m.deltaz)
	}
}

func main() {
	data, _ := ioutil.ReadFile("lol.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	moons := parse(niceData)
	gmoons := applyGravity(moons)
	printMoons(moons)
	fmt.Println()
	printMoons(gmoons)
}
