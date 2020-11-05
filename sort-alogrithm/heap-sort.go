package sort_alogrithm

func HeapSort(nums []int, length int) []int {
	nums = constructHeap(nums, length)
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		downAdjust(nums, 0, i)
	}
	return nums
}

func downAdjust(nums []int, parent int, length int) {
	// 对二叉树的节点进行下沉操作 （小顶堆）
	// 一般用在，排序以及构建堆中
	if parent > length/2 - 1 {
		return
	}
	minChild := minIndexInNums(nums, length, 2*parent+1, 2*parent+2)
	if nums[parent] > nums[minChild] {
		nums[parent], nums[minChild] = nums[minChild], nums[parent]
		downAdjust(nums, minChild, length)
	}

	//for parent < length / 2 {
	//	left := 2 * parent + 1
	//	right := 2 * parent + 2
	//	minIdx := minIndexInNums(nums, length, left, right)
	//	if nums[parent] > nums[minIdx] {
	//		nums[parent], nums[minIdx] = nums[minIdx], nums[parent]
	//		parent = minIdx
	//	} else {
	//		return
	//	}
	//}
}

func upAdjust(nums []int, length int) {
	// 对二叉堆插入的元素进行上浮操作
	// 插入的元素为nums的最后一个
	idx := length - 1
	for idx > 0 {
		parent := (idx - 1) / 2
		if nums[idx] < nums[parent] {
			nums[idx], nums[parent] = nums[parent], nums[idx]
			idx = parent
		} else {
			return
		}
	}
}

func minIndexInNums(nums []int, length, i, j int) int {
	if i > length-1 {
		return j
	}
	if j > length-1 {
		return i
	}
	if nums[i] < nums[j] {
		return i
	} else {
		return j
	}
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
