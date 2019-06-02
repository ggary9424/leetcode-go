package leedcode_go

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	type testCase struct {
		name           string
		input          []int
		expectedResult bool
	}

	testCases := []testCase{
		{
			name:           "[1,2,3,4,5]",
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: true,
		},
		{
			name:           "[5,4,3,2,1]",
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := increasingTriplet(testCase.input)
			if testCase.expectedResult != actualResult {
				t.Errorf("Expected result is %t, but we got %t.", testCase.expectedResult, actualResult)
			}
		})
	}
}
