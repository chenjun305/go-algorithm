package algorithm

/*
* 数组中的第K个最大元素
 */

func findKthLargest(nums []int, k int) int {
	// 构建前K个元素的最小堆
	buildSmallHeap(nums, k)
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			downAjustHeap(nums, k, 0)
		}
	}
	return nums[0]
}

func buildSmallHeap(nums []int, k int) {
	// 遍历所有非叶子结点，下沉调整堆
	for i := k/2 - 1; i >= 0; i-- {
		downAjustHeap(nums, k, i)
	}
}

func downAjustHeap(nums []int, k int, i int) {
	// 用数组存储最小堆，i节点的左孩子为2*i+1，右孩子为2*i+2
	child := 2*i + 1 // 左孩子
	for child < k {
		if child+1 < k && nums[child+1] < nums[child] {
			// 存在右孩子 且 右孩子更小
			child = child + 1
		}
		if nums[i] < nums[child] {
			// 父节点小于孩子节点，说明调整到位了
			break
		}
		nums[i], nums[child] = nums[child], nums[i]
		i, child = child, 2*child+1
	}
}
