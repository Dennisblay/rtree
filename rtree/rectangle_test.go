// rectangle_test.go
package rtree

import "testing"

func TestNewRectangle(t *testing.T) {
	newRect, err := NewRectangle(1, 2, 3, 4)
	if err != nil {
		t.Error(err)
	}
	if newRect.xMin != 1 || newRect.yMin != 2 || newRect.xMax != 3 || newRect.yMax != 4 {
		t.Errorf("Rectangle coordinates are not correct, got: %+v", newRect)
	}
}

func TestRectangleOverlaps(t *testing.T) {
	A, _ := NewRectangle(1, 1, 4, 4)
	B, _ := NewRectangle(2, 2, 5, 5)
	C, _ := NewRectangle(6, 6, 8, 8)
	D, _ := NewRectangle(3, 3, 7, 7)
	E, _ := NewRectangle(10, 10, 12, 12)

	testCases := []struct {
		rect1, rect2 *Rectangle
		expected     bool
	}{
		{A, B, true},  // Overlapping
		{A, C, false}, // Non-overlapping
		{B, D, true},  // Overlapping
		{A, D, true},  // Overlapping
		{C, E, false}, // Non-overlapping
	}

	for _, tc := range testCases {
		result := tc.rect1.Overlaps(tc.rect2)
		if result != tc.expected {
			t.Errorf("Overlap test failed for rectangles %v and %v: got %v, want %v", tc.rect1, tc.rect2, result, tc.expected)
		}
	}
}

func TestRectangleContains(t *testing.T) {
	A, _ := NewRectangle(1, 1, 4, 4)
	B, _ := NewRectangle(2, 2, 5, 5)
	C, _ := NewRectangle(6, 6, 8, 8)
	D, _ := NewRectangle(3, 3, 7, 7)
	E, _ := NewRectangle(10, 10, 12, 12)

	testCases := []struct {
		rect1, rect2 *Rectangle
		expected     bool
	}{
		{rect1: A, rect2: A, expected: true},
		{rect1: A, rect2: B, expected: false},
		{rect1: A, rect2: C, expected: false},
		{rect1: B, rect2: A, expected: false},
		{rect1: B, rect2: B, expected: true},
		{rect1: D, rect2: A, expected: false},
		{rect1: D, rect2: B, expected: false},
		{rect1: D, rect2: C, expected: false},
		{rect1: C, rect2: C, expected: true},
		{rect1: D, rect2: E, expected: false},
	}

	for _, test := range testCases {
		result := test.rect1.Contains(test.rect2)
		if result != test.expected {
			t.Errorf("Expected %v.Contains(%v) to be %v, but got %v", test.rect1, test.rect2, test.expected, result)
		}
	}

}
