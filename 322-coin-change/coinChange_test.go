package _22_coin_change

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	ret := coinChange([]int{2, 5 , 10, 1}, 27)
	fmt.Println(ret)
}