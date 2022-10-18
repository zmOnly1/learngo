package main

import (
	"fmt"
	"learngo2/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.Print()
	root.SetValue(100)

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var bRoot *tree.Node //bRoot == nil
	bRoot.SetValue(200)
	bRoot = &root
	bRoot.SetValue(300)
	bRoot.Print()

	fmt.Println("start traverse...")
	root.Traverse()
}
