package leetcode_go

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	type input struct {
		m int
		n int
	}
	type testCase struct {
		name           string
		input          *input
		expectedResult int
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: &input{
				m: 5,
				n: 7,
			},
			expectedResult: 4,
		},
		{
			name: "Test Case 2",
			input: &input{
				m: 0,
				n: 1,
			},
			expectedResult: 0,
		},
		{
			name: "Test Case 3",
			input: &input{
				m: 2147483646,
				n: 2147483647,
			},
			expectedResult: 2147483646,
		},
		{
			name: "Test Case 4",
			input: &input{
				m: 2131245352,
				n: 2012338413,
			},
			expectedResult: 1879048192,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := rangeBitwiseAnd(testCase.input.m, testCase.input.n)
			if actualResult != testCase.expectedResult {
				t.Errorf(
					"The expected result is %d, but we got %d",
					testCase.expectedResult,
					actualResult,
				)
			}
		})
	}
}
