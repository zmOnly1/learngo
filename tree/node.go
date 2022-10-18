package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting nil value ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()

	root.print()
	root.setValue(100)

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var bRoot *treeNode //bRoot == nil
	bRoot.setValue(200)
	bRoot = &root
	bRoot.setValue(300)
	bRoot.print()

	fmt.Println("start traverse...")
	root.traverse()
}
