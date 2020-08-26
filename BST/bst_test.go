package main

import "testing"

func TestInsert(t *testing.T) {
	// Test if node can insert data correctly
	node := Node{
		val: 10,
	}
	node.insert(5)
	node.insert(15)
	node.insert(17)

	if node.left.val != 5 {
		t.Error("Expected value was 5")
	}
	if node.right.val != 15 {
		t.Error("Expected value was 15")
	}
	if node.right.right.val != 17 {
		t.Error("Expected value was 17")
	}

}

func TestContains(t *testing.T) {
	// Contains returns false if value not found
	var tree Tree
	tree.insert(5)
	tree.insert(15)
	tree.insert(20)
	tree.insert(-1)
	tree.insert(-3)
	tree.insert(12)

	// tree.root.printTree()
	if !tree.root.contains(-3) {
		t.Error("Expected value was True")
	}
	if tree.root.contains(0) {
		t.Error("Expected value was False")
	}
}
