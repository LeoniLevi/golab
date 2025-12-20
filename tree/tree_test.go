package tree

//import	"github.com/LeoniLevi/golab/tree"
import "testing"

func TestAdd_01(t *testing.T) {
	myTree := CreateTree()
	myTree.Add(15)
	myTree.Add(45)
	myTree.Add(12)
	myTree.Add(33)
	
	expected := 33
	got := myTree.root.right.left.val
	if expected != got {
		t.Errorf("AddTree failed: %d != %d", expected, got)
	}

}


func TestAdd_02(t *testing.T) {
	myTree := CreateTree()
	myTree.Add(15)
	myTree.Add(45)
	myTree.Add(12)
	myTree.Add(33)
	
	expected := "((12)15((33)45))"
	got := myTree.root.ToListStr()

	if expected != got {
		t.Errorf("AddTree failed: %s != %s", expected, got)
	}
}

func TestFind(t *testing.T) {

	myTree := CreateTree()
	myTree.Add(15)
	myTree.Add(45)
	myTree.Add(12)
	myTree.Add(33)
	
	{
		var expected *IntNode = nil
		got := FindTreeNode(myTree, 11)
		if expected != got {
			t.Errorf("TestFind failed: Found unexisting value ??!!")
		}
	}
	{
		expected := 12
		got := FindTreeNode(myTree, 12).GetValue()
		if expected != got {
			t.Errorf("TestFind failed: %d != %d", expected, got)
		}
	}
}
