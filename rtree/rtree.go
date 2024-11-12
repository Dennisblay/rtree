package rtree

// RtreeNode
type RTreeNode struct {
	isLeaf     bool
	bbox       *Rectangle
	children   []*RTreeNode
	entries    []*Rectangle
	parent     *RTreeNode
	maxEntries int // Maximum node capacity
}

// RTree
type RTree struct {
	root     *RTreeNode
	maxChild int
}

func (rtree RTree) Insert(r *Rectangle) {

}

func NewRTree(maxChild int) *RTree {
	return &RTree{
		root: &RTreeNode{
			isLeaf:   true,
			bbox:     nil,
			children: nil,
			entries:  []*Rectangle{},
		},
		maxChild: maxChild,
	}
}
