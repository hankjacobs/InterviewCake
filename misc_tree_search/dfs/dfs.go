package main

import (
	"fmt"

	"github.com/hankjacobs/interviewcake/misc_tree_search/common"
)

func main() {
	t := common.NewTestTree()
	dfs(t)
}

func dfs(tree *common.Node) {
	s := common.NewStack()
	s.Push(tree)

	for s.Size() != 0 {
		n := s.Pop()

		fmt.Println(n.Value)
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}
}
