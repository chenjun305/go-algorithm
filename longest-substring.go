package main

// 滑动窗口算法
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	start_index, size, max := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if v, exist := m[s[i]]; exist && v >= start_index {
			start_index = v + 1
		}
		m[s[i]] = i
		size = i - start_index + 1
		if size > max {
			max = size
		}
	}
	return max
}
