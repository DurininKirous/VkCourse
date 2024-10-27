package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var array []int
	fmt.Scan(&n)
	readArray(&array, n)
	first, second := findMinPair(array, n)
	fmt.Println(first, second)
}

func readArray(array *[]int, n int) {
	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		*array = append(*array, num)
	}
}

func findMinPair(array []int, n int) (int, int) {
	var minIndex int
	for i := 1; i < n-1; i++ {
		PairDiff := math.Abs(float64(array[i] - array[i+1]))
		if math.Abs(float64(array[minIndex]-array[minIndex+1])) > PairDiff {
			minIndex = i
		}
	}
	return array[minIndex], array[minIndex+1]
}
