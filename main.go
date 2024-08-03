package main

import (
	"fmt"
)

type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(value int64) {
	node := &Node{Value: value}
	if tree.Root == nil {
		tree.Root = node
		return
	}
	currentNode := tree.Root
	for true {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = node
				break
			}
			currentNode = currentNode.Left
			continue
		}
		if currentNode.Right == nil {
			currentNode.Right = node
			break
		}
		currentNode = currentNode.Right
	}
}

func (tree *BinaryTree) Search(value int64) bool {
	currentNode := tree.Root

	for true {
		if currentNode == nil {
			break
		}

		if currentNode.Value == value {
			return true
		}

		if value < currentNode.Value {
			currentNode = currentNode.Left
			continue
		}

		if value > currentNode.Value {
			currentNode = currentNode.Right
			continue
		}
	}

	return false
}

func main() {
	node := &Node{Value: 100}
	binaryTree := BinaryTree{
		Root: node,
	}
	binaryTree.Insert(50)
	binaryTree.Insert(75)
	binaryTree.Insert(160)

	fmt.Printf("Root: %d \n", binaryTree.Root.Value)
	fmt.Printf("Root left: %d \n", binaryTree.Root.Left.Value)
	fmt.Printf("Root right: %d \n", binaryTree.Root.Right.Value)
	fmt.Printf("Root left right: %d \n", binaryTree.Root.Left.Right.Value)
}
