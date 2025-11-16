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
	tree.PrintNodes(myTree.Root)
	fmt.Println()
	fmt.Println("-----------------------")
	tree.DisplayTreeHorizontally(myTree)

	fmt.Println("===== examineIntTree Completed.")
}
