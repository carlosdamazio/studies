package main

import "fmt"

func main() {
	arr1 := []int{5, 2, 4, 3}
	arr2 := []int{5, 2, 4, 7, 6}
	fmt.Println(intersection(arr1, arr2))
}

func intersection(arr1 []int, arr2 []int) []int {
	var result []int

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				result = append(result, arr1[i])
				break
			}
		}
	}

	return result
}
