package main

/*
* 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
* 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
* 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*
* 解决方案:
* 遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？
* 当考虑完所有天数之时，我们就得到了最好的答案。
* 时间复杂度：O(n)O(n)O(n)，只需要遍历一次。
* 空间复杂度：O(1)O(1)O(1)，只使用了常数个变量。
 */
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
