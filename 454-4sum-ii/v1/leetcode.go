package leetcode_go

func fourSumCount(A []int, B []int, C []int, D []int) int {
	if len(A) == 0 {
		return 0
	}

	count := 0
	countMap := map[int]int{}
	for _, i := range A {
		for _, j := range B {
			if _, ok := countMap[i+j]; ok {
				countMap[i+j]++
			} else {
				countMap[i+j] = 1
			}
		}
	}
	for _, i := range C {
		for _, j := range D {
			if val, ok := countMap[(i+j)*-1]; ok {
				count += val
			}
		}
	}

	return count
}
