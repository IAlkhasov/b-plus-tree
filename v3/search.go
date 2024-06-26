package main

import (
	"errors"
	"fmt"
	"slices"
)

/*
explicitly without the generic `Node`:
a leaf node:
type LeafNode struct {
	data    []int
	parent  *InternalNode // Pointer to parent node for easier navigation
	next    *LeafNode     // Pointer to the next leaf node for range queries
}

an internal node:
type InternalNode struct {
	keys     []int
	children []*InternalNode // For internal nodes, children are other internal nodes
	parent   *InternalNode   // Pointer to parent node for easier navigation
}

NB: this example does not demonstrate sibling or parent pointers.
*/

func (n *Node) Search(key int) (*Node, int, error) {
	idx, found := slices.BinarySearch(n.keys, key)

	if found {
		if len(n.children) == 0 {
			return n, idx, nil
		} else {
			return nil, 0, errors.ErrUnsupported
		}

	}

	if len(n.children) == 0 {
		return n, 0, errors.New("key not found, at leaf containing key")
	}

	return n.children[idx].Search(key)
}

func (n *Node) binSearchExplicit(key int) (*Node, int) {
	if n.kind == LEAF_NODE {
		// If it's a leaf node, return the leaf node and the index where the key would be or is found
		return n, binarySearch(n.keys, key)
	}

	// If it's not a leaf node, recursively binSearchExplicit for the appropriate child node
	i := binarySearch(n.keys, key)
	return n.children[i].binSearchExplicit(key)
}

func binarySearch(arr []int, key int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// see for practice: https://leetcode.com/problems/search-insert-position/description/
	return low // where it would be or should be inserted (for insert)
}

func BasicSearchLeaf(tree *BTree, key int) {
	result, index := tree.root.binSearchExplicit(key)

	fmt.Println(result)
	fmt.Println(index)
}
