package main

import "fmt"

func main() {
	var n int
	var array []int
	var target int
	fmt.Scan(&n)
	readArray(&array, n)
	fmt.Scan(&target)
	fmt.Printf("%d", findTargetIndex(array, n, target))
}

func readArray(array *[]int, n int) {
	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		*array = append(*array, num)
	}
}

func findTargetIndex(array []int, n int, target int) int {
	left := 0
	right := n - 1
	if target > array[n-1] {
		return n
	} else if target < array[0] {
		return 0
	}
	for {
		if left >= right {
			break
		}
		middle := (left + right) / 2
		if array[middle] == target {
			return middle
		} else if right-left == 1 {
			return right
		} else if target > array[middle] {
			left = middle
		} else {
			right = middle
		}
	}
	return -1
}
