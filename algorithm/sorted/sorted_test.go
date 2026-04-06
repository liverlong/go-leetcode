package sorted

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{6, 4, 12, 7, 19, 2, 13, 56, 5}

	QuickSort(arr)

	fmt.Println(arr)
}
