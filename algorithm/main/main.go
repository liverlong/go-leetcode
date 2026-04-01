package main

import (
	"fmt"
	"go-lc/algorithm/sorted"
)

func main() {
	arr := []int{6, 4, 12, 7, 123}

	a := sorted.InsertionSort(arr)

	fmt.Println(a)
}
