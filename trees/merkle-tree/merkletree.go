package merkletree

// Merkle tree info https://www.wikiwand.com/en/Merkle_tree

type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	hash  []byte
}

type MerkleTree struct {
	root *MerkleNode
	size int
}

// function to initialize a new Merkle Node
func NewMerkleNode(left, right *MerkleNode, hash []byte) *MerkleNode {
	newNode := MerkleNode{}

	if left == nil && right == nil {
		newNode.hash = hash
	} else {
		prevHashes := append(left.hash, right.hash...)
		newNode.hash = append(hash, prevHashes...)
	}
	newNode.left = left
	newNode.right = right

	return &newNode
}

// function to insert a new node into the Merkle Tree
func (tree *MerkleTree) Insert(left, right *MerkleNode, hash []byte) {
	newNode := NewMerkleNode(left, right, hash)
	tree.size++

	if tree.root == nil {
		tree.root = newNode
	} else {
		// get the current node
		currentNode := tree.root

		// loop until we find a node that doesn't have a right child
		for currentNode.right != nil {
			currentNode = currentNode.right
		}

		// insert the new node as the right child
		currentNode.right = newNode
	}
}
