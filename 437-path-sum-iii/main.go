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
    cChan := make(chan int)
    nodeCount := traverse(root, func (node *TreeNode) {
        cChan <- calculate(node, sum)
    })

    for i:=0;i<nodeCount;i++ {
        c += <-cChan
    }

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

func traverse(node *TreeNode, callback func (*TreeNode)) int {
    c1, c2 := 0, 0
    if node == nil {
        return 0
    }

    go callback(node)
    if node.Left != nil {
        c1 = traverse(node.Left, callback)
    }
    if node.Right != nil {
        c2 = traverse(node.Right, callback)
    }
    return 1 + c1 + c2
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