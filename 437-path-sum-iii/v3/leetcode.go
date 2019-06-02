package leetcode_go

import (
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeJob struct {
	Callback func(node *TreeNode)
	Node     *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var mutex sync.Mutex
	result := 0
	jobChan := startTraverseWorker(10)
	jobProcess := 0
	nodeCount := traverse(jobChan, root, func(node *TreeNode) {
		mutex.Lock()
		result += calculate(node, sum)
		jobProcess += 1
		mutex.Unlock()
	})

	for jobProcess != nodeCount {
		// do nothing
	}
	return result
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

func startTraverseWorker(number int) chan TreeNodeJob {
	treeNodeJobs := make(chan TreeNodeJob, 10*number)
	for i := 0; i < number; i++ {
		go func() {
			for true {
				j := <-treeNodeJobs
				j.Callback(j.Node)
			}
		}()
	}
	return treeNodeJobs
}

func traverse(jobChan chan TreeNodeJob, node *TreeNode, callback func(*TreeNode)) int {
	c1, c2 := 0, 0
	if node == nil {
		return 0
	}

	jobChan <- TreeNodeJob{
		Callback: callback,
		Node:     node,
	}

	if node.Left != nil {
		c1 = traverse(jobChan, node.Left, callback)
	}
	if node.Right != nil {
		c2 = traverse(jobChan, node.Right, callback)
	}
	return 1 + c1 + c2
}
