package array

//4. 寻找两个有序数组的中位数
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	total := len(nums1) + len(nums2)
	count := 0
	stop := total/2 + 1
	i, j := 0, 0
	heap := make([]int, 0)
	for count < stop {
		if i >= len(nums1) {
			heap = append(heap, nums2[j])
			j++
		} else if j >= len(nums2) {
			heap = append(heap, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			heap = append(heap, nums1[i])
			i++
		} else {
			heap = append(heap, nums2[j])
			j++
		}
		count++
	}
	if total%2 == 0 {
		return float64(heap[len(heap)-1]+heap[len(heap)-2]) / 2
	} else {
		return float64(heap[len(heap)-1])
	}
}
