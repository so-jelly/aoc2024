package maths

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
