package tree

import "fmt"
import "strings"

type IntNode struct {
	val   int
	left  *IntNode
	right *IntNode
}

type IntTree struct {
	root *IntNode
}

func (myTree *IntTree) Add(nval int) {
	addTreeNode(myTree, nval)
}


func (root *IntNode) ToListStr() string {
	str := ""
	if root != nil {
		var sbuilder strings.Builder
		sbuilder.WriteString("(")
		
		sleft := root.left.ToListStr()
		sval := fmt.Sprintf("%d", root.val)
		sright := root.right.ToListStr()

		sbuilder.WriteString(sleft)
		sbuilder.WriteString(sval)
		sbuilder.WriteString(sright)
		sbuilder.WriteString(")")
/*
		str = str + "("
		str = str + root.left.ToListStr()
		str = str + fmt.Sprintf("%d", root.val)
		str = str + root.right.ToListStr()
		str = str + ")"
		*/
		str = sbuilder.String()
	}
	return str
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

func addTreeNode(tree *IntTree, nval int) {
	newNode := IntNode{val: nval, left: nil, right: nil}
	if tree.root == nil {
		tree.root = &newNode
	} else {
		addNewNode(tree.root, &newNode)
	}
}

func FindTreeNode(tree *IntTree, nval int) *IntNode {
	return findNode(tree.root, nval)
}

/*
func RemoveTreeNode(tree *IntTree, nval int) {

}
*/

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

func findNode(root *IntNode, nval int) *IntNode {
	for node := root; node != nil; {
		if node.val == nval {
			return node
		}
		if nval < node.val {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func displayNodeSubtreeHorizontally(root *IntNode, indent int) {
	if root != nil {
		displayNodeSubtreeHorizontally(root.left, indent+3)
		fmt.Printf("%*s%d\n", indent, "", root.val)
		displayNodeSubtreeHorizontally(root.right, indent+3)
	}
}
