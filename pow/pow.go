package pow

// PowX uses recursion to solve x^n
func PowX(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / PowX(x, -n)
	}

	if n == 2 {
		return square(x)
	}

	if n%2 == 0 {
		return PowX(square(x), n/2)
	}

	return x * PowX(x, n-1)
}

func square(x float64) float64 {
	return x * x
}
