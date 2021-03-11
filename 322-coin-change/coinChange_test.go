package _22_coin_change

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	ret := coinChange3([]int{1,2,5}, 11)
	fmt.Println(ret)
}