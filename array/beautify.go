package array

//1248. 统计「优美子数组」
//给你一个整数数组 nums 和一个整数 k。
//
//如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
//
//请返回这个数组中「优美子数组」的数目。
//
//
//
//示例 1：
//
//输入：nums = [1,1,2,1,1], k = 3
//输出：2
//解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
//示例 2：
//
//输入：nums = [2,4,6], k = 1
//输出：0
//解释：数列中不包含任何奇数，所以不存在优美子数组。
//示例 3：
//
//输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
//输出：16

func numberOfSubarrays(nums []int, k int) int {
	count := 0
	n := 0
	heap := make([]int, 0)
	more := make([]int, 0)
	for next := 0; next < len(nums); next++ {
		cur := nums[next]
		if cur%2 == 1 {
			//找到奇数
			n++
			if n <= k {
				heap = append(heap, cur)
			}
		} else {
			if n < k {
				heap = append(heap, cur)
			} else {
				more = append(more, cur)
			}
		}
		if n > k {
			first := 0
			for i, h := range heap {
				if h%2 == 1 {
					first = i
					break
				}
			}
			count += (first + 1) * (len(more) + 1)
			n = k
			if first+1 < len(heap) {
				heap = heap[first+1:]
			} else {
				heap = make([]int, 0)
			}
			heap = append(heap, more...)
			heap = append(heap, cur)
			more = make([]int, 0)
		}
	}
	if n == k {
		first := 0
		for i, h := range heap {
			if h%2 == 1 {
				first = i
				break
			}
		}
		count += (first + 1) * (len(more) + 1)
	}

	return count
}
