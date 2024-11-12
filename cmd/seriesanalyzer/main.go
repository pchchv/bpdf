package main

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
}
