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

func TestIntTree() {

	fmt.Println("===== TestIntTree started...")
	fmt.Println(" ~~ adding items to tree..")

	var tree *IntTree = createTree()
	addTreeNode(tree, 10)
	addTreeNode(tree, 15)

	addTreeNode(tree, 4)
	addTreeNode(tree, 12)
	addTreeNode(tree, 6)
	addTreeNode(tree, 1)
	addTreeNode(tree, 17)
	addTreeNode(tree, 11)
	addTreeNode(tree, 7)

	fmt.Println(" ~~ result tree(after adding):")
	printNodes(tree.root)
	fmt.Println()
	fmt.Println("-----------------------")
	displayBT(tree.root, 0)

	fmt.Println("===== TestIntTree Completed.")

}

func createTree() *IntTree {
	tree := IntTree{root: nil}
	return &tree
	//return new IntTree {root: nil}
}

func addTreeNode(tree *IntTree, nval int) {
	newNode := IntNode{val: nval, left: nil, right: nil}
	if tree.root == nil {
		tree.root = &newNode
	} else {
		addNewNode(tree.root, &newNode)
	}
}

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

func printNodes(root *IntNode) {
	fmt.Print("(")
	if root != nil {
		if root.left != nil {
			printNodes(root.left)
		}
		fmt.Print(root.val)
		if root.right != nil {
			printNodes(root.right)
		}
	}
	fmt.Print(")")
}

func displayBT(root *IntNode, indent int) {
	if root != nil {
		displayBT(root.left, indent+3)
		fmt.Printf("%*s%d\n", indent, "", root.val)
		displayBT(root.right, indent+3)
	}
}
