package medium

import (
	"fmt"
	"testing"
)

func TestCoin(t *testing.T) {
	coins := []int{1, 2, 5}

	fmt.Println(coinChange(coins, 11))
}
