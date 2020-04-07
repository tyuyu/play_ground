package stack

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//
//
//
//进阶：
//
//你能在线性时间复杂度内解决此题吗？
//
//
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7

func maxSlidingWindow(nums []int, k int) []int {

	res := make([]int, 0)

	h := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if len(h) == 0 || num < h[0] {
			h = append(h, num)
			if len(h) > k {
				mi := 1
				for i := 2; i < len(h); i++ {
					if h[i] > h[mi] {
						mi = i
					}
				}
				h = h[mi:]
			}
		} else {
			h = []int{num}
		}
		if i+1 >= k {
			res = append(res, h[0])
		}
	}
	return res
}
