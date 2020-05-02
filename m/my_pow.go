package m

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return float64(1) / myPow(x, -n)
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return tmp * tmp * x
	}
}
