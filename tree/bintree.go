package tree

import "fmt"

type IntNode struct {
	val   int
	left  *IntNode
	right *IntNode
}

type IntTree struct {
	root *IntNode
}

func (myTree *IntTree) GetRoot() *IntNode {
	return myTree.root
}

func (node *IntNode) GetValue() int {
	return node.val
}

func (node *IntNode) GetLeft() *IntNode {
	return node.left
}

func (node *IntNode) GetRight() *IntNode {
	return node.right
}

func CreateTree() *IntTree {
	tree := IntTree{root: nil}
	return &tree
	//return new IntTree {root: nil}
}

func AddTreeNode(tree *IntTree, nval int) {
	newNode := IntNode{val: nval, left: nil, right: nil}
	if tree.root == nil {
		tree.root = &newNode
	} else {
		addNewNode(tree.root, &newNode)
	}
}

func DisplayTreeHorizontally(myTree *IntTree) {
	displayNodeSubtreeHorizontally(myTree.root, 0)
}

//-------------------

func addNewNode(parent *IntNode, newnode *IntNode) {
	if newnode.val <= parent.val {
		if parent.left == nil {
			parent.left = newnode
		} else {
			addNewNode(parent.left, newnode)
		}
	} else {
		if parent.right == nil {
			parent.right = newnode
		} else {
			addNewNode(parent.right, newnode)
		}
	}
}

func displayNodeSubtreeHorizontally(root *IntNode, indent int) {
	if root != nil {
		displayNodeSubtreeHorizontally(root.left, indent+3)
		fmt.Printf("%*s%d\n", indent, "", root.val)
		displayNodeSubtreeHorizontally(root.right, indent+3)
	}
}
