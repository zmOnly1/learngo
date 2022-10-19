package main

import (
	"fmt"
	"learngo2/tree"
)

type myTreeNode struct {
	*tree.Node //Embedding, child implement cannot assign to base type
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	leftNode := myTreeNode{myNode.Left}
	rightNode := myTreeNode{myNode.Right}

	leftNode.postOrder()
	rightNode.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("start traverse...")
	root.Traverse()
	root.Node.Traverse()

	fmt.Println("start post traverse...")
	root.postOrder()
}
