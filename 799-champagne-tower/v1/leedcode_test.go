package leedcode_go

import "testing"

func TestChampagneTower(t *testing.T) {
	type input struct {
		poured     int
		queryRow   int
		queryGlass int
	}
	type testCase struct {
		name           string
		input          input
		expectedResult float64
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: input{
				poured:     1,
				queryRow:   1,
				queryGlass: 1,
			},
			expectedResult: 0,
		},
		{
			name: "Test Case 2",
			input: input{
				poured:     2,
				queryRow:   1,
				queryGlass: 1,
			},
			expectedResult: 0.5,
		},
		{
			name: "Test Case 3",
			input: input{
				poured:     2,
				queryRow:   2,
				queryGlass: 1,
			},
			expectedResult: 0,
		},
		{
			name: "Test Case 4",
			input: input{
				poured:     6,
				queryRow:   2,
				queryGlass: 1,
			},
			expectedResult: 1,
		},
		{
			name: "Test Case 5",
			input: input{
				poured:     6,
				queryRow:   2,
				queryGlass: 0,
			},
			expectedResult: 0.75,
		},
		{
			name: "Test Case 6",
			input: input{
				poured:     5050,
				queryRow:   55,
				queryGlass: 55,
			},
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := champagneTower(testCase.input.poured, testCase.input.queryRow, testCase.input.queryGlass)
			if testCase.expectedResult != actualResult {
				t.Errorf("Expected result is %f, but we got %f.", testCase.expectedResult, actualResult)
			}
		})
	}
}
