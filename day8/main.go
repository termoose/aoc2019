package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const Width = 25
const Height = 6

type layer struct {
	pixels []int
}

func getXY(count int) (int, int) {
	x := count % Width
	y := count / Width
	return x, y
}

func calcLayers(input string) []layer {
	var output []layer
	layerSize := Width * Height
	inputSize := len(input)
	nrLayers := inputSize / layerSize

	for i := 0; i < nrLayers; i++ {
		low := i * layerSize
		high := (i + 1) * layerSize

		var pixels []int
		for _, elem := range input[low:high] {
			pixel, _ := strconv.Atoi(string(elem))
			pixels = append(pixels, pixel)
		}

		elem := layer{
			pixels: pixels,
		}

		output = append(output, elem)
	}

	return output
}

func countElems(elem int, layer layer) int {
	count := 0
	for _, pixel := range layer.pixels {
		if pixel == elem {
			count++
		}
	}

	return count
}

func minElemsLayer(elem int, layers []layer) layer {
	minLayer := layer{}
	minCount := len(layers[0].pixels) + 1

	for _, layer := range layers {
		count := countElems(elem, layer)
		if count < minCount {
			minCount = count
			minLayer = layer
		}
	}

	return minLayer
}

func addLayers(first, second layer) layer {
	var result layer

	for i, firstPixel := range first.pixels {
		secondPixel := second.pixels[i]

		if firstPixel == 2 {
			result.pixels = append(result.pixels, secondPixel)
		} else {
			result.pixels = append(result.pixels, firstPixel)
		}
	}

	return result
}

func combineLayers(layers []layer) layer {
	var first layer
	var result layer

	for i := 0; i < len(layers) - 1; i++ {
		if i == 0 {
			first = layers[i]
		} else {
			first = result
		}

		second := layers[i + 1]

		result = addLayers(first, second)
	}

	return result
}

func printLayer(layer layer) {
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			pixel := layer.pixels[h * Width + w]
			if pixel == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}

		fmt.Println()
	}
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	niceData := strings.TrimSuffix(string(data), "\n")

	layers := calcLayers(niceData)
	maxLayer := minElemsLayer(0, layers)
	count := countElems(0, maxLayer)
	fmt.Printf("Max layer: %d Count 0: %d\n", maxLayer, count)

	countOnes := countElems(1, maxLayer)
	countTwos := countElems(2, maxLayer)
	fmt.Printf("Ones x Twos = %d\n", countOnes * countTwos)

	combined := combineLayers(layers)
	printLayer(combined)
}
