package main

import (
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

var Tree = &TreeNode{
	&TreeNode{
		nil, nil, 9,
	}, &TreeNode{
		nil, nil, 8,
	},
	10,
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	var que []*TreeNode
	que = append(que, root)
	for len(que) != 0 {
		head := que[0]
		if head.left != nil {
			que = append(que, head.left)
		}
		if head.right != nil {
			que = append(que, head.right)
		}
		fmt.Println(head.value)
		que = que[1:]
	}
}

func main() {
	//bfs(Tree)
	var str1 []string
	var str2 []string
	str1 = append(str1, )
}
