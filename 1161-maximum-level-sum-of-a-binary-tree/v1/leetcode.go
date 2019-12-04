package leetcode_go

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	var levelCountMap map[int]int
	levelCountMap = make(map[int]int)

	traverse(root, 1, func(node *TreeNode, level int) {
		if val, ok := levelCountMap[level]; ok {
			levelCountMap[level] = val + node.Val
		} else {
			levelCountMap[level] = node.Val
		}
	})

	return findMaxValueOfMapAndReturnKey(levelCountMap)
}

func traverse(node *TreeNode, level int, callback func(*TreeNode, int)) {
	if node == nil {
		return
	}
	callback(node, level)
	if node.Left != nil {
		traverse(node.Left, level+1, callback)
	}
	if node.Right != nil {
		traverse(node.Right, level+1, callback)
	}
}

func findMaxValueOfMapAndReturnKey(m map[int]int) int {
	levelOfMaxValue := 1
	max := m[levelOfMaxValue]
	for k, v := range m {
		if v > max {
			max = v
			levelOfMaxValue = k
		}
	}
	return levelOfMaxValue
}
