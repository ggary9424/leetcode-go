package leetcode_go

import (
	"testing"
)

func TestPartition(t *testing.T) {
	s := []int{3, 1, 4, 1, 5}
	v := 2
	expected := 2
	actual := findPairs(s, v)
	if actual != expected {
		t.Errorf(
			"Input is %v and %d. Expected value is %d, but we got %d",
			s, v, expected, actual,
		)
		return
	}

	s = []int{1, 2, 3, 4, 5}
	v = 1
	expected = 4
	actual = findPairs(s, v)
	if actual != expected {
		t.Errorf(
			"Input is %v and %d. Expected value is %d, but we got %d",
			s, v, expected, actual,
		)
		return
	}

	s = []int{1, 3, 1, 5, 4}
	v = 0
	expected = 1
	actual = findPairs(s, v)
	if actual != expected {
		t.Errorf(
			"Input is %v and %d. Expected value is %d, but we got %d",
			s, v, expected, actual,
		)
		return
	}

	s = []int{1, 1, 1, 1, 1}
	v = 0
	expected = 1
	actual = findPairs(s, v)
	if actual != expected {
		t.Errorf(
			"Input is %v and %d. Expected value is %d, but we got %d",
			s, v, expected, actual,
		)
		return
	}

	s = []int{}
	v = 2
	expected = 0
	actual = findPairs(s, v)
	if actual != expected {
		t.Errorf(
			"Input is %v and %d. Expected value is %d, but we got %d",
			s, v, expected, actual,
		)
		return
	}
}
