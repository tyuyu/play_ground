package array

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

func merge(nums1 []int, m int, nums2 []int, n int) {

	var tmp []int
	tmp = append(tmp, nums1[0:m]...)

	index, i, j := 0, 0, 0
	for index < len(nums1) {
		if i >= len(tmp) {
			nums1[index] = nums2[j]
			j++
		} else if j >= len(nums2) {
			nums1[index] = tmp[i]
			i++
		} else {
			a1, a2 := tmp[i], nums2[j]
			if a1 <= a2 {
				nums1[index] = a1
				i++
			} else {
				nums1[index] = a2
				j++
			}
		}
		index++
	}
	return
}

func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return deep(nums, k)
}

func deep(nums []int, k int) [][]int {
	ans := make([][]int, 0)
	if k == 1 {
		for _, num := range nums {
			ans = append(ans, []int{num})
		}
		return ans
	}

	for i, num := range nums {
		for _, xx := range deep(nums[i+1:], k-1) {
			ans = append(ans, append([]int{num}, xx...))
		}
	}

	return ans
}
