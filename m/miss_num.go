package m

func firstMissingPositive(nums []int) int {

	bits := make([]byte, len(nums)+2)
	for _, num := range nums {
		if num < 0 {
			continue
		}
		if num > len(nums)+1 {
			continue
		}
		bits[num] = 1
	}
	for i := 1; i < len(bits); i++ {
		if bits[i] == 0 {
			return i
		}
	}
	return 1
}
