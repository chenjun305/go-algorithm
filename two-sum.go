package main

import "fmt"

/*
* 给定一个整数数组 nums 和一个整数目标值 target，
* 请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
*
* 思路及算法
* 创建一个哈希表，对于每一个 num，我们首先查询哈希表中是否存在 target - num，然后将 num 插入到哈希表中，即可保证不会让 num 和自己匹配。
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
