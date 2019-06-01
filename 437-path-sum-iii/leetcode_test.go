package leetcode_go

import (
    "testing"
)

func TestPathSum(t *testing.T) {
    tree := &TreeNode{
        Val: 10,
        Left: &TreeNode{
            Val: 5,
            Left: &TreeNode{
                Val: 3,
                Left: &TreeNode{
                    Val: 3,
                    Left: nil,
                    Right: nil,
                },
                Right: &TreeNode{
                    Val: -2,
                    Left: nil,
                    Right: nil,
                },
            },
            Right: &TreeNode{
                Val: 2,
                Left: nil,
                Right: &TreeNode{
                    Val: 1,
                    Left: nil,
                    Right: nil,
                },
            },
        },
        Right: &TreeNode{
            Val: -3,
            Left: nil,
            Right: &TreeNode{
                Val: 11,
                Left: nil,
                Right: nil,
            },
        },
    }
    
    result := pathSum(tree, 8)
    if result != 3 {
        t.Errorf("The result should be 3, but we got %d", result)
    }
}