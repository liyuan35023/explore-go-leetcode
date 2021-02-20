package __median_two_sorted_array

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T)  {
	k := findMedianSortedArrays2([]int{1,2}, []int{3,4})
	fmt.Println(k)
}
