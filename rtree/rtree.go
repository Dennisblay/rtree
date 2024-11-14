package rtree

import (
	"fmt"
	"math"
)

// RtreeNode
type Node struct {
	isLeaf   bool
	bbox     BBox
	children []*Node
	entries  []*Rectangle
	parent   *Node
}

type BBox interface {
	Overlaps(other *Rectangle) bool
	Contains(other *Rectangle) bool
	Area() float64
	Extend(other *Rectangle)
	Union(other *Rectangle) *Rectangle
	Equals(other *Rectangle) bool
	Distance(other *Rectangle) float64
}

func (n *Node) ResizeBBox(r BBox) {
	if n.bbox == nil {
		n.bbox = r
	} else {
		n.bbox.Extend(r.(*Rectangle))
	}
}

func (n *Node) PushEntry(r *Rectangle) {
	n.entries = append(n.entries, r)
	n.ResizeBBox(r)
}

func (n *Node) PushChild(node *Node) {
	n.children = append(n.children, node)
	n.ResizeBBox(node.bbox)
}

// RTree
type RTree struct {
	root       *Node
	maxEntries int // Maximum node capacity
}

func (rtree *RTree) Insert(r *Rectangle) {
	leaf := rtree.chooseSubtree(r)

	leaf.PushEntry(r)

	rtree.adjustTree(leaf)
}

func (rtree *RTree) splitNode(node *Node) (*Node, *Node) {
	firstSeed, secondSeed := rtree.chooseSplitSeeds(node)

	node1 := &Node{
		isLeaf:   node.isLeaf,
		children: nil,
		entries:  nil,
	}

	node2 := &Node{
		isLeaf:   node.isLeaf,
		children: nil,
		entries:  nil,
	}

	// Push seeds and Resize nodes to fit bounds
	node1.PushEntry(firstSeed)
	node2.PushEntry(secondSeed)

	// Distribute the remaining entries to the two nodes
	for _, entry := range node.entries {
		if entry != firstSeed && entry != secondSeed {
			enlargement1, _ := bboxEnlargementArea(node1.bbox, entry)
			enlargement2, _ := bboxEnlargementArea(node2.bbox, entry)
			if enlargement1 < enlargement2 {
				node1.PushEntry(entry)
			} else {
				node2.PushEntry(entry)
			}
		}
	}
	return node1, node2

}

func (rtree *RTree) chooseSplitSeeds(leaf *Node) (*Rectangle, *Rectangle) {
	var firstSeed, secondSeed *Rectangle
	var maxDistance = -1.0

	// Find two nodes that are farthest apart
	for i := 0; i < len(leaf.entries); i++ {
		for j := i + 1; j < len(leaf.entries); j++ {
			distance := leaf.entries[i].Distance(leaf.entries[j])
			if distance > maxDistance {
				maxDistance = distance
				firstSeed = leaf.entries[i]
				secondSeed = leaf.entries[j]
			}
		}
	}
	return firstSeed, secondSeed
}

func (rtree *RTree) adjustTree(node *Node) {

	if len(node.entries) > rtree.maxEntries {
		rtree.splitNode(node)
	}

	if node.parent == nil {
		newRoot := &Node{
			isLeaf:   false,
			bbox:     nil,
			children: []*Node{},
			entries:  []*Rectangle{},
		}

		newRoot.PushChild(node)
	}

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
			enlargement, err := bboxEnlargementArea(child.bbox, rect)
			if err != nil {
				panic(err)
			}

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
func bboxEnlargementArea(bbox, rect BBox) (float64, error) {
	bboxRect, ok := bbox.(*Rectangle)
	if !ok {
		return 0, fmt.Errorf("bbox is not of type *Rectangle")
	}

	rectRect, ok := rect.(*Rectangle)
	if !ok {
		return 0, fmt.Errorf("rect is not of type *Rectangle")
	}

	// Early return if bbox already contains rect
	if bboxRect.Contains(rectRect) {
		return 0, nil
	}

	// Calculate the dimensions of the new bounding box
	newBBox := bboxRect.Union(rectRect)

	// Calculate the area of the expanded bounding box and subtract the original area
	enlargedArea := newBBox.Area() - bboxRect.Area()
	return enlargedArea, nil
}
