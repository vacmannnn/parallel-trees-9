package simpleTree

import (
	"strconv"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (n *Node) insert(val int) {
	if val <= n.val {
		if n.left == nil {
			n.left = &Node{val: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{val: val}
	} else {
		t.root.insert(val)
	}
}

func (t *Tree) Find(key int) bool {
	var keyFound bool
	if t.root == nil {
		keyFound = false
	} else if t.root.val != key {
		keyFound = t.root.find(key)
	} else {
		keyFound = true
	}
	return keyFound
}
func (n *Node) find(key int) bool {
	var keyFound bool
	if n == nil {
		return false
	}
	if key == n.val {
		return true
	}
	if n.val > key {
		keyFound = n.left.find(key)
	} else {
		keyFound = n.right.find(key)
	}
	return keyFound
}

// Remove removes the Item with key `key` from the tree
func (t *Tree) Remove(key int) {
	remove(t.root, key)
}

// internal recursive function to remove an item
func remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if val < node.val {
		node.left = remove(node.left, val)
		return node
	}
	if val > node.val {
		node.right = remove(node.right, val)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftMostRightSide := node.right
	for {
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}
	node.val = leftMostRightSide.val
	node.right = remove(node.right, node.val)
	return node
}

func String(node *Node) string {
	if node == nil {
		return ""
	}
	s := String(node.left) + " " + strconv.Itoa(node.val) + " " + String(node.right)
	return s
}

func (t *Tree) String() string {
	return String(t.root)
}

func (t *Tree) IsValid() bool {
	return isValid(t.root)
}

func isValid(root *Node) bool {
	return RecValidate(root, nil, nil)
}

func RecValidate(n, min, max *Node) bool {
	if n == nil {
		return true
	}
	if min != nil && n.val <= min.val {
		return false
	}
	if max != nil && n.val >= max.val {
		return false
	}
	return RecValidate(n.left, min, n) && RecValidate(n.right, n, max)
}
