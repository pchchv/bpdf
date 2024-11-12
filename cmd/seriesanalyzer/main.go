package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mergeArrays(a []float64, b []float64) (merged []float64) {
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}

	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return
}

func MergeFloat64(arr []float64) []float64 {
	if len(arr) < 2 {
		return arr
	}

	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]

	return mergeArrays(MergeFloat64(left), MergeFloat64(right))
}

func main() {
	bytes, err := os.ReadFile("docs/assets/text/benchmark.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	stringContent := string(bytes)

	var sum float64
	var values []float64
	lines := strings.Split(stringContent, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		values = append(values, value)
		sum += value
	}

	values = MergeFloat64(values)

	fmt.Printf("min: %f, max: %f, avg: %f", values[0], values[len(values)-1], sum/float64(len(values)))
}
