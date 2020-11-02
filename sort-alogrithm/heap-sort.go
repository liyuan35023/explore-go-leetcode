package sort_alogrithm

func HeapSort(nums []int, left int, right int) {
	
}

func downAdjust(nums []int, parent int, length int) {
	// 对二叉树的节点进行下沉操作 （小顶堆）
	// 一般用在，排序以及构建堆中
	if parent > (length-2) / 2 {
		return
	}
	minChild := minIndexInNums(nums, 2*parent+1, 2*parent+2)
	if nums[parent] > nums[minChild] {
		nums[parent], nums[minChild] = nums[minChild], nums[parent]
		downAdjust(nums, minChild, length)
	}
}

func minIndexInNums(nums []int, i, j int) int {
	if i > len(nums) - 1 {
		return j
	}
	if j > len(nums) - 1 {
		return i
	}
	if nums[i] < nums[j] {
		return i
	} else {
		return j
	}
}

func upAdjust(nums []int, length int) {
	// 对二叉堆插入的元素进行上浮操作
	// 插入的元素为nums的最后一个


}

func constructHeap(nums []int, length int) []int {
	// 构建堆(最小堆)
	// 从非叶子节点开始进行下沉操作，(顺序是自底向上)
	lastNotleaf := (length - 2) / 2
	for i := lastNotleaf; i >= 0; i-- {
		downAdjust(nums, i, length)
	}
	return nums
}

