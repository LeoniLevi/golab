package tree

import "fmt"

type IntNode struct {
	Val   int
	Left  *IntNode
	Right *IntNode
}

type IntTree struct {
	Root *IntNode
}

func CreateTree() *IntTree {
	tree := IntTree{Root: nil}
	return &tree
	//return new IntTree {root: nil}
}

func AddTreeNode(tree *IntTree, nval int) {
	newNode := IntNode{Val: nval, Left: nil, Right: nil}
	if tree.Root == nil {
		tree.Root = &newNode
	} else {
		addNewNode(tree.Root, &newNode)
	}
}

func PrintNodes(root *IntNode) {
	fmt.Print("(")
	if root != nil {
		if root.Left != nil {
			PrintNodes(root.Left)
		}
		fmt.Print(root.Val)
		if root.Right != nil {
			PrintNodes(root.Right)
		}
	}
	fmt.Print(")")
}

func DisplayTreeHorizontally(myTree *IntTree) {
	displayNodeSubtreeHorizontally(myTree.Root, 0)
}

//-------------------

func addNewNode(parent *IntNode, newnode *IntNode) {
	if newnode.Val <= parent.Val {
		if parent.Left == nil {
			parent.Left = newnode
		} else {
			addNewNode(parent.Left, newnode)
		}
	} else {
		if parent.Right == nil {
			parent.Right = newnode
		} else {
			addNewNode(parent.Right, newnode)
		}
	}
}

func displayNodeSubtreeHorizontally(root *IntNode, indent int) {
	if root != nil {
		displayNodeSubtreeHorizontally(root.Left, indent+3)
		fmt.Printf("%*s%d\n", indent, "", root.Val)
		displayNodeSubtreeHorizontally(root.Right, indent+3)
	}
}
