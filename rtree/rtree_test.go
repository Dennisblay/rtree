package rtree

import "testing"

func TestNewRTree(t *testing.T) {
	maxChild := 4
	rt := NewRTree(maxChild)

	if rt.root == nil {
		t.Error("RTree 'root' should not be nil")
	}

	if rt.root.isLeaf != true {
		t.Error("RTree should be a leaf node initially")
	}

	if len(rt.root.entries) != 0 {
		t.Error("RTree should have no entries initially")
	}

	if rt.maxChild != maxChild {
		t.Errorf("Expected maxChild to be %v, but got: %v", maxChild, rt.maxChild)
	}

}
