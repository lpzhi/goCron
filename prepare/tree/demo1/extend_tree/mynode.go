package extend_tree

import (
	"fmt"
	"goCron/prepare/tree/demo1/tree"
)

//组合的方式来扩展包
type MyNode struct {
	Node *tree.Node
}

//前序遍历

func (node *MyNode) PreorderTraversal() {

	if node == nil || node.Node == nil {
		return
	}

	left := MyNode{node.Node.Left}
	right := MyNode{node.Node.Right}
	fmt.Println(node.Node.DataValue)
	left.PreorderTraversal()
	right.PreorderTraversal()

}
