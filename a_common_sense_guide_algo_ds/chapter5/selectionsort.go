package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 3}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	var lowestNumberIndex int

	for i := 0; i < len(arr) - 1; i++ {
		lowestNumberIndex = i
		for j := i ; j < len(arr); j++ {
			if arr[j] < arr[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}
		if lowestNumberIndex != i {
			arr[i], arr[lowestNumberIndex] = arr[lowestNumberIndex], arr[i]
		}
	}

	return arr
}