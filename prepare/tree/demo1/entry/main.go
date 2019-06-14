package main

import (
	"fmt"
	"goCron/prepare/tree/demo1/extend_tree"
	"goCron/prepare/tree/demo1/tree"
)

type q []int

func main() {
	root := tree.Node{DataValue: 3}
	root.Left = new(tree.Node)
	root.Right = new(tree.Node)
	root.Left.DataValue = 5
	root.Right.DataValue = 9
	root.Traverse()
	myNode := extend_tree.MyNode{&root}

	myNode.PreorderTraversal()

	q1 := append(q{2}, 1)

	fmt.Println(q1)

}
