package number

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func majorityElement(nums []int) int {
	half := len(nums) / 2
	major := 0
	c := 0
	for _, num := range nums {
		if c == 0 {
			major = num
			c++
			continue
		}
		if c > half {
			return major
		}
		if num == major {
			c++
		} else {
			c--
		}
	}
	return major
}
