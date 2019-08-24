package leetcode_go

import "testing"

func TestTrimBST(t *testing.T) {
	type input struct {
		tree            *TreeNode
		lowestBoundary  int
		highestBoundary int
	}

	type testCase struct {
		name           string
		input          input
		expectedResult *TreeNode
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: input{
				tree:            nil,
				lowestBoundary:  0,
				highestBoundary: 100,
			},
			expectedResult: nil,
		},
		{
			name: "Test Case 2",
			input: input{
				tree: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  0,
						Left: nil,
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				lowestBoundary:  1,
				highestBoundary: 3,
			},
			expectedResult: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		},
		{
			name: "Test Case 3",
			input: input{
				tree: &TreeNode{
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
				lowestBoundary:  1,
				highestBoundary: 2,
			},
			expectedResult: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult := trimBST(
				testCase.input.tree,
				testCase.input.lowestBoundary,
				testCase.input.highestBoundary,
			)
			if !checkTwoTreeHaveSameVal(actualResult, testCase.expectedResult) {
				t.Errorf("Answer of \"%s\" is wrong", testCase.name)
			}
		})
	}
}

func checkTwoTreeHaveSameVal(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 != nil && tree2 != nil {
		if tree1.Val != tree2.Val {
			return false
		}

		return checkTwoTreeHaveSameVal(tree1.Left, tree2.Left) && checkTwoTreeHaveSameVal(tree1.Right, tree2.Right)
	}
	if tree1 == nil && tree2 == nil {
		return true
	}
	return false
}
