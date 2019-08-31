package leetcode_go

import "testing"

func TestAverageOfLevels(t *testing.T) {
	type testCase struct {
		name           string
		treeNode       *TreeNode
		expectedResult []float64
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			treeNode: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			expectedResult: []float64{1, 1},
		},
		{
			name: "Test Case 2",
			treeNode: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expectedResult: []float64{3, 14.5, 11},
		},
		{
			name: "Test Case 3",
			treeNode: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 17,
						Left: &TreeNode{
							Val: 22,
							Left: &TreeNode{
								Val:   3,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			expectedResult: []float64{3, 14.5, 16, 22, 3},
		},
		{
			name:           "Test Case 4",
			treeNode:       nil,
			expectedResult: []float64{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := averageOfLevels(testCase.treeNode)
			if len(actualResult) != len(testCase.expectedResult) {
				t.Errorf("The length of actualResult is not same with length of expectedResult.\n")
			}

			for level, result := range actualResult {
				if result != testCase.expectedResult[level] {
					t.Errorf(
						"Expected average of level '%d' is %f, but we got %f.\n",
						level,
						testCase.expectedResult[level],
						result,
					)
				}
			}
		})
	}
}
