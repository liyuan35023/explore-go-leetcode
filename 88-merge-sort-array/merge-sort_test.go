package _8_merge_sort_array

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{4,5,6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}


