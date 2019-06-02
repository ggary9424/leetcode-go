package leedcode_go

const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	n1 := MaxInt
	n2 := MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] <= n1 {
			n1 = nums[i]
		} else if nums[i] <= n2 {
			n2 = nums[i]
		} else {
			return true
		}
	}
	return false
}
