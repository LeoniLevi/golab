package main

import (
	"fmt"

	"github.com/LeoniLevi/golab/tree"
)

func main() {
	examineIntTree()
}

func examineIntTree() {

	fmt.Println("===== examineIntTree started...")
	fmt.Println(" ~~ adding items to tree..")

	var myTree = tree.CreateTree()
	tree.AddTreeNode(myTree, 15)

	tree.AddTreeNode(myTree, 4)
	tree.AddTreeNode(myTree, 12)
	tree.AddTreeNode(myTree, 10)
	tree.AddTreeNode(myTree, 6)
	tree.AddTreeNode(myTree, 1)
	tree.AddTreeNode(myTree, 17)
	tree.AddTreeNode(myTree, 11)
	tree.AddTreeNode(myTree, 7)
	tree.AddTreeNode(myTree, 16)

	fmt.Println(" ~~ result tree(after adding):")
	printNodes(myTree.GetRoot())
	fmt.Println()
	fmt.Println("-----------------------")
	tree.DisplayTreeHorizontally(myTree)

	fmt.Println("===== examineIntTree Completed.")
}

func printNodes(root *tree.IntNode) {
	fmt.Print("(")
	if root != nil {
		if root.GetLeft() != nil {
			printNodes(root.GetLeft())
		}
		fmt.Print(root.GetValue())
		if root.GetRight() != nil {
			printNodes(root.GetRight())
		}
	}
	fmt.Print(")")
}
