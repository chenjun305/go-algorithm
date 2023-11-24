package main

import "fmt"

/*
* 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
* 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
* 思路及解法
* 使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
* 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
 */
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			if right > left {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 4, 5, 0}
	moveZeroes(nums1)
	fmt.Println(nums1)
	nums2 := []int{0, 0, 1, 2, 0, 3}
	moveZeroes(nums2)
	fmt.Println(nums2)
}
