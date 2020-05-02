package array

import "sort"

//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func amerge(intervals [][]int) [][]int {

	res := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
		}
		last := res[len(res)-1]
		if interval[0] < last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			res = append(res, interval)
		}
	}

	return res
}

//57. 插入区间
//给出一个无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//示例 1:
//
//输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出: [[1,5],[6,9]]
//示例 2:
//
//输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出: [[1,2],[3,10],[12,16]]
//解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

func insert(intervals [][]int, newInterval []int) [][]int {

	res := make([][]int, 0)

	tmp := [][]int{newInterval}
	last := make([][]int, 0)
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			last = append(last, interval)
		} else {
			tmp = append(tmp, interval)
		}
	}

	middle := amerge(tmp)
	res = append(append(res, middle...), last...)
	return res
}
