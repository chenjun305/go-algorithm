package main

func longestConsecutive(nums []int) int {
	// 构建1个哈希存储所有数，方便后续查找
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	longest := 0
	for num := range m {
		// 如果比当前num更小的也存在，则跳过
		if !m[num-1] {
			curNum := num
			curLenth := 1
			for m[curNum+1] {
				curLenth++
				curNum++
			}
			if curLenth > longest {
				longest = curLenth
			}
		}
	}
	return longest
}
