package tree

type Node struct {
	DataValue interface{}
	Left      *Node
	Right     *Node
}
