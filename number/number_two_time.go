package number

//面试题56 - I. 数组中数字出现的次数
//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
//
//
//示例 1：
//
//输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
//示例 2：
//
//输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2]
func singleNumbers(nums []int) []int {

	res := 0
	for _, num := range nums {
		res ^= num
	}

	div := 1
	for res&div == 0 {
		div = div << 1
	}

	o := []int{0, 0}
	for _, num := range nums {
		if num&div == 0 {
			o[0] = o[0] ^ num
		} else {
			o[1] = o[1] ^ num
		}
	}
	return o
}
