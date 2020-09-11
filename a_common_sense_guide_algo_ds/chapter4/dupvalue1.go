package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(hasDuplicatedValues(arr))
}

func hasDuplicatedValues(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				return true
			}
		}
	}

	return false
}