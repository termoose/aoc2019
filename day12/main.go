package main

import (
	"crypto/sha256"
	_ "crypto/sha256"
	_ "encoding/gob"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type universe struct {
	moons [4]moon
}

func (u *universe) hash() string {
	mStr := fmt.Sprintf("%v", u)
	result := fmt.Sprintf("%x", sha256.Sum256([]byte(mStr)))

	return result
}

type moon struct {
	x, y, z int
	vx, vy, vz int
}

func (m *moon) hash() string {
	mStr := fmt.Sprintf("%v", m)
	result := fmt.Sprintf("%x", sha256.Sum256([]byte(mStr)))

	return result
}

func (m *moon) energy() int {
	potential := int(math.Abs(float64(m.x)) + math.Abs(float64(m.y)) + math.Abs(float64(m.z)))
	kinetic := int(math.Abs(float64(m.vx)) + math.Abs(float64(m.vy)) + math.Abs(float64(m.vz)))

	return potential * kinetic
}

func (m *moon) gravitate(other moon) {
	if m.x > other.x {
		m.vx--
	} else if m.x < other.x {
		m.vx++
	}

	if m.y > other.y {
		m.vy--
	} else if m.y < other.y {
		m.vy++
	}

	if m.z > other.z {
		m.vz--
	} else if m.z < other.z {
		m.vz++
	}
}

func (m *moon) move() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func parse(data string) universe {
	var result universe

	for i, line := range strings.Split(data, "\n") {
		elem := moon{}

		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &elem.x, &elem.y, &elem.z)
		result.moons[i] = elem
		//result.moons = append(result.moons, elem)
	}

	return result
}

func applyGravity(u universe) universe {
	var result universe

	for i, m1 := range u.moons {
		for _, m2 := range u.moons {
			m1.gravitate(m2)
		}

		result.moons[i] = m1
		//result.moons = append(result.moons, m1)
	}

	return result
}

func iterate(moons universe) universe {
	gravityMoons := applyGravity(moons)
	var result universe

	for i, moon := range gravityMoons.moons {
		moon.move()
		result.moons[i] = moon
		//result.moons = append(result.moons, moon)
	}

	return result
}

func printMoons(u universe) {
	totalEnergy := 0
	for _, m := range u.moons {
		totalEnergy += m.energy()

		fmt.Printf("Energy: %d <x=%2d, y=%2d, z=%2d> v: <%2d, %2d, %2d>\n",
			m.energy(), m.x, m.y, m.z,
			m.vx, m.vy, m.vz)
		fmt.Printf("Hash: %s\n", m.hash())
	}

	fmt.Printf("Total energy: %d\n", totalEnergy)
}

func main() {
	data, _ := ioutil.ReadFile("lol.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	moons := parse(niceData)

	printMoons(moons)
	fmt.Println()

	visited := make(map[string]bool)
	iterMoons := moons
	iterations := 0
	for {
		iterMoons = iterate(iterMoons)
		hash := iterMoons.hash()
		if visited[hash] {
			fmt.Printf("Iters %d\n", iterations)
			return
		}

		visited[hash] = true
		iterations++
	}

	//iterMoons := moons
	//for i := 0; i < 1000; i++ {
	//	iterMoons = iterate(iterMoons)
	//	printMoons(iterMoons)
	//	fmt.Println()
	//}
}
