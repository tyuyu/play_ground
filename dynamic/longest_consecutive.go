package dynamic

//最长连续序列
//给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为 O(n)。
//
//示例:
//
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}

	max := 0
	for _, num := range nums {
		count := 1
		left, right := num-1, num+1
		_, ok := m[left]
		for ok {
			delete(m, left)
			left--
			count++
			_, ok = m[left]
		}
		_, ok = m[right]
		for ok {
			delete(m, right)
			right++
			count++
			_, ok = m[right]
		}
		if count > max {
			max = count
		}
	}

	return max
}
