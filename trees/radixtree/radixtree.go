package main

import (
	"fmt"
)

type Radix16Node struct {
	children []*Radix16Node
	value    interface{}
}

func NewRadix16Node() *Radix16Node {
	return &Radix16Node{
		children: make([]*Radix16Node, 16),
		value:    nil,
	}
}

type Radix16Tree struct {
	root *Radix16Node
	size int
}

func NewRadix16Tree() *Radix16Tree {
	return &Radix16Tree{
		root: NewRadix16Node(),
		size: 0,
	}
}

func (t *Radix16Tree) Insert(key string, value interface{}) {
	node := t.root
	for _, char := range key {
		index := hexToDecimal(char)
		if node.children[index] == nil {
			node.children[index] = NewRadix16Node()
		}
		node = node.children[index]
	}
	node.value = value
	t.size++
}

func (t *Radix16Tree) Search(key string) (interface{}, bool) {
	node := t.root
	for _, char := range key {
		index := hexToDecimal(char)
		if node.children[index] == nil {
			return nil, false
		}
		node = node.children[index]
	}
	return node.value, true
}

/*
*
  - hexToDecimal converts a hexademical character to a decimal representation in the range 0-15
*/
func hexToDecimal(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	}
	if char >= 'a' && char <= 'f' {
		return int(char - 'a' + 10)
	}
	if char >= 'A' && char <= 'F' {
		return int(char - 'A' + 10)
	}
	panic(fmt.Sprintf("invalid character: %c", char))
}

func (t *Radix16Tree) Print() {
	printNode(t.root, "")
}

func printNode(node *Radix16Node, prefix string) {
	if node.value != nil {
		fmt.Printf("%s --> %v\n", prefix, node.value)
	}
	for i, child := range node.children {
		if child != nil {
			printNode(child, fmt.Sprintf("%s%c", prefix, decimalToHex(i)))
		}
	}
}
func decimalToHex(index int) rune {
	if index >= 0 && index <= 9 {
		return rune('0' + index)
	}
	return rune('a' + index - 10)
}

func main() {
	tree := NewRadix16Tree()
	tree.Insert("0AB12", "example1")
	tree.Insert("0AB99", "example2")
	tree.Insert("12345", "example3")
	tree.Print()
	fmt.Println(tree.Search("0AB12"))
	fmt.Println(tree.Search("0AB99"))
	fmt.Println(tree.Search("12345"))
	fmt.Println(tree.Search("ABCD"))
}
