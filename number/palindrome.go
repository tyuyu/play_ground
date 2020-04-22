package number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	half := 0
	for x > half {
		last := x % 10
		half = half*10 + last
		x = x / 10
	}
	if half == x || x/10 == half {
		return true
	}
	return false
}
