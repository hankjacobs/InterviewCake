package main

import (
	"fmt"

	"github.com/hankjacobs/interviewcake/misc_tree_search/common"
)

func main() {
	t := common.NewTestTree()
	bfs(t)
}

func bfs(tree *common.Node) {
	s := common.NewQueue()
	s.Enqueue(tree)

	for s.Size() != 0 {
		n := s.Dequeue()

		if n.Left != nil {
			s.Enqueue(n.Left)
		}
		fmt.Println(n.Value)
		if n.Right != nil {
			s.Enqueue(n.Right)
		}
	}
}
