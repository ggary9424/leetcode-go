package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }

    c := 0
    traverse(root, func (node *TreeNode) {
        c += calculate(node, sum)
    })
    return c
}

func calculate(root *TreeNode, sum int) int {
    v := sum - root.Val
    c1, c2 := 0, 0
    if root.Left != nil {
        c1 = calculate(root.Left, v)
    }
    if root.Right != nil {
        c2 = calculate(root.Right, v)
    }
    if v == 0 {
        return 1 + c1 + c2
    }
    return c1 + c2
}

func traverse(node *TreeNode, callback func (*TreeNode)) {
    if node == nil {
        return
    }
    callback(node)
    if node.Left != nil {
        traverse(node.Left, callback)
    }
    if node.Right != nil {
        traverse(node.Right, callback)
    }
}

func main() {
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
    println(pathSum(tree, 8))
}