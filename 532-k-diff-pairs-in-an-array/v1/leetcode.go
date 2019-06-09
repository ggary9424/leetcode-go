package leetcode_go

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	count := 0
	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 1
		} else {
			m[n]++
		}
	}
	if k == 0 {
		for _, v := range m {
			if v > 1 {
				count++
			}
		}
	} else {
		for num, _ := range m {
			if _, ok := m[num+k]; ok {
				count++
			}
			if _, ok := m[num-k]; ok {
				count++
			}
			delete(m, num)
		}
	}
	return count
}
