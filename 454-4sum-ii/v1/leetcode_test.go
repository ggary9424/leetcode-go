package leetcode_go

import "testing"

func TestFourSumCount(t *testing.T) {
	type input struct {
		a []int
		b []int
		c []int
		d []int
	}

	type testCase struct {
		name           string
		input          input
		expectedResult int
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: input{
				a: []int{1, 2},
				b: []int{-2, -1},
				c: []int{-1, 2},
				d: []int{0, 2},
			},
			expectedResult: 2,
		},
		{
			name: "Test Case 2",
			input: input{
				a: []int{1, 2, 3, 4, 5, 6},
				b: []int{-2, -1, -2, -4, -8, -1},
				c: []int{-1, 2, 3, 2, 1, -3},
				d: []int{0, 2, -2, -6, 0, 0},
			},
			expectedResult: 105,
		},
		{
			name: "Test Case 3",
			input: input{
				a: []int{},
				b: []int{},
				c: []int{},
				d: []int{},
			},
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := fourSumCount(
				testCase.input.a,
				testCase.input.b,
				testCase.input.c,
				testCase.input.d,
			)
			if actualResult != testCase.expectedResult {
				t.Errorf(
					"The actual result is %d, but we expected the result is %d",
					actualResult,
					testCase.expectedResult,
				)
			}
		})
	}
}
