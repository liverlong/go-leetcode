package medium

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}

	res := permute(nums)

	for _, data := range res {
		fmt.Println(data)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
