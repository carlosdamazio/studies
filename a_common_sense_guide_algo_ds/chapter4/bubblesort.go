package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 3}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	arrIndex := len(arr) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < arrIndex; i++ {
			if arr[i] >  arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		arrIndex -= 1
	}

	return arr
}
