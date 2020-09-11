package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(hasDuplicatedValues2(arr))
}

func hasDuplicatedValues2(arr []int) bool {
	cacheValues := make([]int, findMax(arr)+1)

	for i := 0; i < len(arr); i++ {
		if cacheValues[arr[i]] == 1 {
			return true
		} else {
			cacheValues[arr[i]] = 1
		}
	}

	return false
}

func findMax(arr []int) int {
	max := 0

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}