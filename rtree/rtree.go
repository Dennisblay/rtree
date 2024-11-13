package rtree

import (
	"math"
)

// RtreeNode
type Node struct {
	isLeaf   bool
	bbox     *Rectangle
	children []*Node
	entries  []*Rectangle
	parent   *Node
}

func (n *Node) ResizeBBox(r *Rectangle) {
	n.bbox.Extend(r)
}

// RTree
type RTree struct {
	root       *Node
	maxEntries int // Maximum node capacity
}

func (rtree *RTree) Insert(r *Rectangle) {
	leaf := rtree.chooseSubtree(r)

	leaf.entries = append(leaf.entries, r)
	leaf.ResizeBBox(r)

	if len(leaf.entries) > rtree.maxEntries {
		rtree.splitNode(leaf)
	}

	rtree.adjustTree(leaf)
}

func (rtree *RTree) splitNode(node *Node) {
	firstSeed, secondSeed := rtree.chooseSplitSeeds(node)

	// isLeaf   bool
	// bbox     *Rectangle
	// children []*Node
	// entries  []*Rectangle
	// parent   *Node

	node1 := &Node{
		isLeaf:   node.isLeaf,
		children: nil,
		entries:  []*Rectangle{firstSeed},
	}

	node2 := &Node{
		isLeaf:   node.isLeaf,
		children: nil,
		entries:  []*Rectangle{secondSeed},
	}

	node1.ResizeBBox(firstSeed)
	node2.ResizeBBox(secondSeed)

	// Distribute the remaining entries to the two nodes
	for _, entry := range node.entries {
		if entry != firstSeed || entry != secondSeed {
			enlargement1, enlargement2 := bboxEnlargementArea(node1.bbox, entry), bboxEnlargementArea(node2.bbox, entry)
			if enlargement1 < enlargement2 {
				node1.entries = append(node1.entries, entry)
			} else {
				node1.entries = append(node1.entries, entry)
			}
		}
	}

}

func (rtree *RTree) chooseSplitSeeds(leaf *Node) (*Rectangle, *Rectangle) {
	var firstSeed, secondSeed *Rectangle
	var maxDistance float64

	// Find two nodes that are farthest apart
	for i := 0; i < len(leaf.entries); i++ {
		for j := i + 1; j < len(leaf.entries); j++ {
			distance := leaf.entries[i].Distance(leaf.entries[j])
			if distance < maxDistance {
				maxDistance = distance
				firstSeed = leaf.entries[i]
				secondSeed = leaf.entries[j]
			}
		}
	}
	return firstSeed, secondSeed
}

func (rtree *RTree) adjustTree(leaf *Node) {

}

// chooseSubtree traverses the RTree from root to an appropriate leaf node for insertion.
func (rtree *RTree) chooseSubtree(rect *Rectangle) *Node {
	current := rtree.root

	// Traverse until reaching a leaf node
	for !current.isLeaf {
		minEnlargementArea := math.MaxFloat64
		var bestChild *Node

		// Select the child node that requires the least area enlargement
		for _, child := range current.children {
			enlargement := bboxEnlargementArea(child.bbox, rect)

			// Choose the smallest enlargement; if tied, pick the node with the smaller area
			if enlargement < minEnlargementArea || (enlargement == minEnlargementArea && child.bbox.Area() < bestChild.bbox.Area()) {
				minEnlargementArea = enlargement
				bestChild = child
			}
		}
		current = bestChild
	}

	return current
}

// NewRTree initializes an RTree with a specified maximum number of child nodes per node.
func NewRTree(maxEntries int) *RTree {
	return &RTree{
		root: &Node{
			isLeaf:   true,
			bbox:     nil,
			children: nil,
			entries:  []*Rectangle{},
		},
		maxEntries: maxEntries,
	}
}

// bboxEnlargementArea calculates the additional area required to expand bbox to include rect.
func bboxEnlargementArea(bbox, rect *Rectangle) float64 {
	// Early return if bbox already contains rect
	if bbox.Contains(rect) {
		return 0
	}

	// Calculate the dimensions of the new bounding box
	newBboxMaxX := math.Max(bbox.maxX, rect.maxX)
	newBboxMinX := math.Min(bbox.minX, rect.minX)
	newBboxMaxY := math.Max(bbox.maxY, rect.maxY)
	newBboxMinY := math.Min(bbox.minY, rect.minY)

	// Calculate the area of the expanded bounding box and subtract the original area
	newArea := (newBboxMaxX - newBboxMinX) * (newBboxMaxY - newBboxMinY)
	return newArea - bbox.Area()
}
