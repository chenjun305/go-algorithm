package main

import "fmt"

/*
* 给定一个整数数组 nums 和一个整数目标值 target，
* 请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
 */
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if p, ok := m[target-num]; ok {
			return []int{p, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{1, 3, 6, 10}
	target := 9
	fmt.Println("input: nums=", nums, ", target=", target)
	fmt.Println("result:", twoSum(nums, target))
}
