package leetcode_go

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, H int) int {
	if len(piles) == 0 {
		return 0
	}
	max, _ := maxInSlice(piles)
	for i := 1; i <= max; i++ {
		sum := 0
		for _, pile := range piles {
			sum += int(math.Ceil(float64(pile) / float64(i)))
		}
		if sum <= H {
			return i
		}
	}
	return 0
}

func maxInSlice(s []int) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("Slice cannot be empty.")
	}
	m := s[0]
	for i, e := range s {
		if i == 0 || e > m {
			m = e
		}
	}
	return m, nil
}
