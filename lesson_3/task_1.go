package main

import (
	"fmt"
)

func main() {
	var n int
	var array []int
	fmt.Scan(&n)
	var index1 int = n
	readArray(&array, n)
	findFistZero(array, n, &index1)
	replaceElements(&array, n, index1)
	print(array, n)
}

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func print(array []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(array[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}

func findFistZero(array []int, n int, index1 *int) {
	for i := 0; i < n; i++ {
		if array[i] == 0 {
			*index1 = i
			break
		}
	}
}

func readArray(array *[]int, n int) {
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		*array = append(*array, num)
	}
}

func replaceElements(array *[]int, n int, index1 int) {
	for i := index1 + 1; i < n; i++ {
		if (*array)[i] != 0 {
			swap(&(*array)[i], &(*array)[index1])
			index1++
		}
	}
}
