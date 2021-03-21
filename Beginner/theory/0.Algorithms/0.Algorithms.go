package main

import (
	"fmt"
)

func main() {
	//linerAlgorithms()
	//fmt.Println("____________")
	//binaryAlgorithms()
	fmt.Println("____________")
	selectionSort()
}

//
//
func linerAlgorithms() {
	a := []int{3, 4, 5, 2, 6, 7, 10}
	var x int
	fmt.Scanln(&x)
	i, count := linerSearch(a, x)
	fmt.Println("index", i, "number of iteration", count)
}
func linerSearch(arr []int, item int) (int, count int) {

	for i := 0; i < len(arr); i++ {

		count++
		if arr[i] == item {
			fmt.Println(arr[i], "is found")
			return i, count
		}
	}
	fmt.Println("not found")
	return 0, count
}

//
//
func binaryAlgorithms() {
	a := []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14, 15}
	var x int
	fmt.Scanln(&x)
	i, count := binarySearch(a, x)
	fmt.Println("index", i, "number of iteration", count)
}
func binarySearch(arr []int, item int) (index int, count int) {
	start := 0
	end := len(arr)
	found := false

	position := -1
	for found == false && start <= end {
		count++
		var middle = int((start + end) / 2)
		switch {
		case arr[len(arr)-1] < item || item < arr[0]:
			return 0, count
		case arr[middle] == item:
			found = true
			position = middle
			return position, count

		case item < arr[middle]:
			end = middle - 1
		default:
			start = middle + 1

		}

	}
	return 0, count
}

//
//
func selectionSort() {
	//3, 4, 6, 7, 8, 9, 10.2, -2, 3.5,
	arr := []float64{3, 4, -2, 3.5, -17}
	count := 0
	for i, val := range arr {
		minEl := val
		minIndex := 0
		for j := i + 1; j < len(arr); j++ {
			if minEl > arr[j] {

				minEl = arr[j]
				minIndex = j
			}
		}

		if val > minEl {
			arr[i] = minEl
			arr[minIndex] = val
		}
		fmt.Println(minEl)
	}
	fmt.Println(count, arr)
}
