/*Directions
1) Implement the Node class to create
a binary search Tree.  The constructor
should initialize values 'data', 'left',
and 'right'.
2) Implement the 'insert' method for the
Node class.  Insert should accept an argument
'data', then create an insert a new Node
at the appropriate location in the Tree.
3) Implement the 'contains' method for the Node
class.  Contains should accept a 'data' argument
and return the Node in the Tree with the same value.
*/
package main

import (
	"fmt"
	"math/rand"
)

// Node definition ...
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Tree definition
type Tree struct {
	root *Node
}

func (t *Tree) insert(data int) {
	// If there is no data in root Node then insert one
	if t.root == nil {
		t.root = &Node{
			val: data,
		}
	} else {
		// else insert a Node to the Tree
		(t.root).insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.val {
		if n.left == nil {
			n.left = &Node{
				val: data,
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{
				val: data,
			}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *Node) contains(data int) bool {
	var res bool
	if n == nil {
		res = false
	} else if n.val == data {
		res = true
	} else if n.val > data {
		res = n.left.contains(data)
	} else if n.val < data {
		res = n.right.contains(data)
	}
	return res
}

func (n *Node) printTree() {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	n.left.printTree()
	n.right.printTree()
}

func main() {
	fmt.Println("--------- Implementing BST -----------")
	fmt.Println("Inserting Data into BST")
	var t Tree
	for i := 0; i < 10; i++ {
		t.insert(rand.Intn(100))
	}
	fmt.Println("Lets see how our Tree Looks")
	t.root.printTree()
	fmt.Println("Let's see is val 40 exists in our Tree or not")
	res := t.root.contains(40)
	fmt.Println("Does number 40 exists in our Tree ?", res)
}
