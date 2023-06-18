package main

import (
	"fmt"
	"strings"
)

type Node struct {
	children []*Node
	value    interface{}
}

func print() {
	// Create the nodes
	nodeA := &Node{value: "A"}
	nodeB := &Node{value: "B"}
	nodeC := &Node{value: "C"}
	nodeD := &Node{value: "D"}
	nodeE := &Node{value: "E"}
	nodeF := &Node{value: "F"}
	nodeG := &Node{value: "G"}
	nodeH := &Node{value: "H"}

	// Link the nodes to form a tree
	nodeA.children = []*Node{nodeB, nodeC, nodeD}
	nodeB.children = []*Node{nodeE, nodeF}
	nodeC.children = []*Node{nodeG}
	nodeD.children = []*Node{nodeH}

	// Print the tree structure
	printTree(nodeA, 0)
}

func printTree(node *Node, depth int) {
	if node == nil {
		return
	}

	// Print the current node value with indentation based on depth
	fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.value)

	// Recursively print the children
	for _, child := range node.children {
		printTree(child, depth+1)
	}
}
