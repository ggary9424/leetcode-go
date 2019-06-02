package leetcode_go

import "testing"

func TestLengthOfLIS(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedResult int
	}

	testCases := []testCase{
		{
			name:           "[]",
			input:          []int{},
			expectedResult: 0,
		},
		{
			name:           "[10,9,2,5,3,7,101,18]",
			input:          []int{10, 9, 2, 5, 3, 7, 101, 18},
			expectedResult: 4,
		},
		{
			name:           "[5,10,3,2,5,6,3,7,10,18,14,67,32]",
			input:          []int{5, 10, 3, 2, 5, 6, 3, 7, 10, 18, 14, 67, 32},
			expectedResult: 7,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := lengthOfLIS(testCase.input)
			if testCase.expectedResult != actualResult {
				t.Errorf("Expected result is %d, but we got %d.", testCase.expectedResult, actualResult)
			}
		})
	}
}
