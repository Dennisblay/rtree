package rtree

import (
	"math"
)

// RtreeNode
type Node struct {
	isLeaf     bool
	bbox       *Rectangle
	children   []*Node
	entries    []*Rectangle
	parent     *Node
	maxEntries int // Maximum node capacity
}

// RTree
type RTree struct {
	root     *Node
	maxChild int
}

func (rtree *RTree) Insert(r *Rectangle) {

}

// chooseSubtree traverses the RTree from root to an appropriate leaf node for insertion.
func (rtree *RTree) chooseSubtree(rect *Rectangle) *Node {
	if rtree.root == nil {
		return nil // Optionally, return an error here or handle this case in the calling function.
	}
	current := rtree.root

	// Traverse until reaching a leaf node
	for !current.isLeaf {
		minEnlargementArea := math.MaxFloat64
		var bestChild *Node

		// Select the child node that requires the least area enlargement
		for _, child := range current.children {
			enlargement := bboxEnlargementArea(child.bbox, rect)

			// Choose the smallest enlargement; if tied, pick the node with the larger area
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
func NewRTree(maxChild int) *RTree {
	return &RTree{
		root: &Node{
			isLeaf:   true,
			bbox:     nil,
			children: nil,
			entries:  []*Rectangle{},
		},
		maxChild: maxChild,
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
