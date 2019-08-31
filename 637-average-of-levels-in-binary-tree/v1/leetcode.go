package leetcode_go

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type record struct {
	Total int
	Count int
}

func averageOfLevels(root *TreeNode) []float64 {
	recordMap := map[int]record{}
	traverseTreeAndRecord(root, 0, recordMap)
	result := []float64{}
	for i := 0; i < len(recordMap); i++ {
		result = append(result, float64(recordMap[i].Total)/float64(recordMap[i].Count))
	}
	return result
}

func traverseTreeAndRecord(node *TreeNode, curretLevel int, recordMap map[int]record) {
	if node == nil {
		return
	}

	if _, ok := recordMap[curretLevel]; ok {
		recordMap[curretLevel] = record{
			Total: recordMap[curretLevel].Total + node.Val,
			Count: recordMap[curretLevel].Count + 1,
		}
	} else {
		recordMap[curretLevel] = record{
			Total: node.Val,
			Count: 1,
		}
	}

	traverseTreeAndRecord(node.Left, curretLevel+1, recordMap)
	traverseTreeAndRecord(node.Right, curretLevel+1, recordMap)
}
