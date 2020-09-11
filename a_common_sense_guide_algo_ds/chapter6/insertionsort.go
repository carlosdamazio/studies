package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int  {
	for i := 1; i < len(arr); i++ {
		tempValue := arr[i]
		position := i - 1

		for position >= 0 {
			if arr[position] > tempValue {
				arr[position + 1] = arr[position]
			} else {
				break
			}
			position -= 1
		}

		arr[position + 1] = tempValue
	}

	return arr
}