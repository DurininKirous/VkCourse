package main

import "fmt"

func main() {
	var n int
	var array []int
	var target int
	fmt.Scan(&n)
	readArray(&array, n)
	fmt.Scan(&target)
	removeElement(&array, &n, target)
	printArray(array, n)
}

func readArray(array *[]int, n int) {
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		*array = append(*array, num)
	}
}

func removeElement(array *[]int, n *int, target int) {
	for i := 0; i < *n; i++ {
		if (*array)[i] == target {
			copy((*array)[i:], (*array)[i+1:])
			(*array)[*n-1] = 0
			(*array) = (*array)[:*n-1]
			*n -= 1
			i -= 1
		}
	}
}

func printArray(array []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print((array)[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}
