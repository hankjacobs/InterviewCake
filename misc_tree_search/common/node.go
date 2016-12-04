package common

// Node node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode new node
func NewNode(value int, left *Node, right *Node) *Node {
	return &Node{value, left, right}
}

//             4
//          2    5
//        1   3    6
//          7  10
//         8  11
//        9
// NewTestTree new
func NewTestTree() *Node {
	n1 := NewNode(1, nil, nil)
	n2 := NewNode(2, nil, nil)
	n3 := NewNode(3, nil, nil)
	n4 := NewNode(4, nil, nil)
	n5 := NewNode(5, nil, nil)
	n6 := NewNode(6, nil, nil)
	n7 := NewNode(7, nil, nil)
	n8 := NewNode(8, nil, nil)
	n9 := NewNode(9, nil, nil)
	n10 := NewNode(10, nil, nil)
	n11 := NewNode(11, nil, nil)

	n2.Left = n1
	n2.Right = n3

	n3.Left = n7
	n3.Right = n10

	n7.Left = n8

	n8.Left = n9

	n10.Left = n11

	n4.Left = n2

	n4.Right = n5
	n5.Right = n6

	return n4
}
