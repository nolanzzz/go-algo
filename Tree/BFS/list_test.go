package BFS

import (
	"container/list"
	"fmt"
	"testing"
)

func TestLevel(t *testing.T) {
	leftLeft := TreeNode{Val: 4, Left: nil, Right: nil}
	leftRight := TreeNode{Val: 5, Left: nil, Right: nil}
	left := TreeNode{Val: 2, Left: &leftLeft, Right: &leftRight}
	right := TreeNode{Val: 3, Left: nil, Right: nil}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	fmt.Println(levelOrder(&root))
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := list.New()
	q.PushBack(root)
	for q.Len() != 0 {
		size := q.Len()
		levelRes := make([]int, size)
		for i := 0; i < size; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			levelRes[i] = node.Val
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		res = append(res, levelRes)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
