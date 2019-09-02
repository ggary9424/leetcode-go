package leetcode_go

func rangeBitwiseAnd(m int, n int) int {
	var shiftCount uint32 = 0

	for m != 0 && m != n {
		m = m >> 1
		n = n >> 1
		shiftCount++
	}

	return m << shiftCount
}
