package sorted

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{6, 4, 12, 7, 123}

	MergeSortBU(arr)

	fmt.Println(arr)
}
