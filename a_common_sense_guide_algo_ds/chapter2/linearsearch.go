package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 10}
	fmt.Println(linearSearch(arr1, 6))
}

func linearSearch(arr []int, val int) bool {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == val {
			return true
		} else if val < arr[i] {
			return false
		}
	}

	return false
}