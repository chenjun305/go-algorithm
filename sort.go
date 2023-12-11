package main

import "fmt"

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums[start]
	left, right := start, end
	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		if left < right {
			// 交换
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[start], nums[left] = nums[left], pivot

	// 递归
	quicksort(nums, start, left-1)
	quicksort(nums, left+1, end)
}

func main() {
	// nums := []int{3, 1, 4, 5, 2, 6, 8, 3}
	nums := []int{}
	fmt.Scan(&nums)
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
