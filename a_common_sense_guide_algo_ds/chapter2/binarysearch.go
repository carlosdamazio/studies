package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 10, 20, 30, 44, 47, 49}
	fmt.Println(binarySearch(arr1, 22))
}

func binarySearch(arr []int, val int) bool {
	lowerBound := 0
	upperBound := len(arr) - 1

	for lowerBound <= upperBound {
		midPoint := (lowerBound + upperBound) / 2
		valAtMid := arr[midPoint]

		if val == valAtMid {
			return true
		} else if val < valAtMid {
			upperBound = midPoint - 1
		} else {
			lowerBound = midPoint + 1
		}
	}

	return false
}
