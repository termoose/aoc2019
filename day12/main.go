package main

import (
	"crypto/md5"
	"crypto/sha256"
	_ "encoding/gob"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const Bodies = 4

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type universe struct {
	moons [Bodies]moon
}

func (u *universe) hash() string {
	mStr := fmt.Sprintf("%v", u)
	result := fmt.Sprintf("%x", md5.Sum([]byte(mStr)))

	return result
}

func (u *universe) energy() int {
	total := 0
	for _, m := range u.moons {
		total += m.energy()
	}
	return total
}

func (u *universe) same(other universe) bool {
	for i, b := range u.moons {
		if b.x != other.moons[i].x || b.y != other.moons[i].y ||
			b.z != other.moons[i].z {
			return false
		}
	}

	return true
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

func (m *moon) gravitateX(other *moon) {
	if m.x > other.x {
		m.vx--
	} else if m.x < other.x {
		m.vx++
	}
}

func (m *moon) gravitateY(other *moon) {
	if m.y > other.y {
		m.vy--
	} else if m.y < other.y {
		m.vy++
	}
}

func (m *moon) gravitateZ(other *moon) {
	if m.z > other.z {
		m.vz--
	} else if m.z < other.z {
		m.vz++
	}
}

func (m *moon) gravitate(other *moon) {
	m.gravitateX(other)
	m.gravitateY(other)
	m.gravitateZ(other)
}

func (m *moon) moveX() {
	m.x += m.vx
}

func (m *moon) moveY() {
	m.y += m.vy
}

func (m *moon) moveZ() {
	m.z += m.vz
}

func (m *moon) move() {
	m.moveX()
	m.moveY()
	m.moveZ()
}

func parse(data string) universe {
	var result universe

	for i, line := range strings.Split(data, "\n") {
		elem := moon{}

		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &elem.x, &elem.y, &elem.z)
		result.moons[i] = elem
	}

	return result
}

func applyGravityX(u *universe) {
	for i, _ := range u.moons {
		for _, m2 := range u.moons {
			u.moons[i].gravitateX(&m2)
		}
	}
}

func applyGravityY(u *universe) {
	for i, _ := range u.moons {
		for _, m2 := range u.moons {
			u.moons[i].gravitateY(&m2)
		}
	}
}

func applyGravityZ(u *universe) {
	for i, _ := range u.moons {
		for _, m2 := range u.moons {
			u.moons[i].gravitateZ(&m2)
		}
	}
}

func applyGravity(u *universe) {
	for i, _ := range u.moons {
		for _, m2 := range u.moons {
			u.moons[i].gravitate(&m2)
		}
	}
}

func iterateX(moons *universe) {
	applyGravityX(moons)

	for i, _ := range moons.moons {
		moons.moons[i].moveX()
	}
}

func iterateY(moons *universe) {
	applyGravityY(moons)

	for i, _ := range moons.moons {
		moons.moons[i].moveY()
	}
}

func iterateZ(moons *universe) {
	applyGravityZ(moons)

	for i, _ := range moons.moons {
		moons.moons[i].moveZ()
	}
}

func iterate(moons *universe) {
	applyGravity(moons)

	for i, _ := range moons.moons {
		moons.moons[i].move()
	}
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
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	moons := parse(niceData)

	printMoons(moons)
	fmt.Println()

	xchan := make(chan int)
	ychan := make(chan int)
	zchan := make(chan int)

	go func() {
		xmoons := moons
		iterations := 0
		for {
			iterateX(&xmoons)

			if iterations % 1000 == 0 {
				//fmt.Printf("X iters %d\n", iterations)
			}

			iterations++
			//if xmoons.same(moons) {
			if xmoons == moons {
				xchan <- iterations
				return
			}

		}
	}()

	go func() {
		ymoons := moons
		iterations := 0
		for {
			iterateY(&ymoons)

			iterations++
			//if ymoons.same(moons) {
			if ymoons == moons {
				ychan <- iterations

				return
			}

		}
	}()

	go func() {
		zmoons := moons
		iterations := 0
		for {
			iterateZ(&zmoons)

			iterations++
			//if zmoons.same(moons) {
			if zmoons == moons {
				zchan <- iterations
				return
			}

		}
	}()

	xCycle := 0
	yCycle := 0
	zCycle := 0
	//xs := make(map[int64]bool)
	//ys := make(map[int64]bool)
	//zs := make(map[int64]bool)
	//var xMatch []int64
	//var yMatch []int64
	//var zMatch []int64

	for {
		select {
		case x := <-xchan:
			xCycle = x
			//xs[x] = true
			//
			//if ys[x] && zs[x] {
			fmt.Printf("X Match %d\n", x)
			//	return
			//}
		case y := <- ychan:
			//fmt.Printf("Y Match: %d\n", y)
			yCycle = y
			//ys[y] = true
			//if xs[y] && zs[y] {
			fmt.Printf("Y Match %d\n", y)
			//	return
			//}
			//yMatch = append(yMatch, y)
		case z := <- zchan:
			zCycle = z
			//zs[z] = true
			//if ys[z] && xs[z] {
			fmt.Printf("Z Match %d\n", z)
			//	return
			//}
			//fmt.Printf("Z Match: %d\n", z)
			//zMatch = append(zMatch, z)
		}

		if xCycle != 0 && yCycle != 0 && zCycle != 0 {
			fmt.Printf("%d %d %d\n", xCycle, yCycle, zCycle)
			fmt.Printf("LCM: %d\n", LCM(xCycle, yCycle, zCycle))
			return
		}
	}

	//firstState := moons.hash()
	//visited := make(map[string]bool)
	//iterMoons := moons
	//iterations := 0
	//for {
	//	iterateX(&iterMoons)
		//hashState := iterMoons.hash()

		//if visited[hash] {
		//	fmt.Printf("Iters %d\n", iterations)
		//	return
		//}
	//	if iterMoons == moons {
	//		xMatch = append(xMatch, iterMoons)
	//		fmt.Printf("Universe: %t\nIters: %d\n", iterMoons, iterations)
	//	}
	//	if iterations % 1000000 == 0 {
	//		fmt.Printf("Iters: %d\n", iterations)
	//	}

		//visited[hash] = true
	//	iterations++
	//}
}
