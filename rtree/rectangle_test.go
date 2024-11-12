// rectangle_test.go
package rtree

import "testing"

func TestNewRectangle(t *testing.T) {
	newRect, err := NewRectangle(1, 2, 3, 4)
	if err != nil {
		t.Error(err)
	}
	if newRect.minX != 1 || newRect.minY != 2 || newRect.maxX != 3 || newRect.maxY != 4 {
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

	for _, tc := range testCases {
		result := tc.rect1.Contains(tc.rect2)
		if result != tc.expected {
			t.Errorf("Expected %v.Contains(%v) to be %v, but got %v", tc.rect1, tc.rect2, tc.expected, result)
		}
	}

}

func TestRectangleArea(t *testing.T) {
	A, _ := NewRectangle(1, 1, 4, 4)
	B, _ := NewRectangle(2, 2, 5, 5)
	C, _ := NewRectangle(6, 6, 8, 8)
	D, _ := NewRectangle(3, 3, 7, 7)
	E, _ := NewRectangle(10, 10, 12, 12)

	testCases := []struct {
		rect         *Rectangle
		expectedArea float64
	}{
		{rect: A, expectedArea: 9},
		{rect: B, expectedArea: 9},
		{rect: C, expectedArea: 4},
		{rect: D, expectedArea: 16},
		{rect: E, expectedArea: 4},
	}

	for _, tc := range testCases {
		result := tc.rect.Area()
		if result != tc.expectedArea {
			t.Errorf("Expected Area of rectangle: %+v, to be: %v, but got: %v", tc.rect, tc.expectedArea, result)
		}
	}
}

func TestRectangleEquals(t *testing.T) {
	A, _ := NewRectangle(1, 1, 4, 4)
	B, _ := NewRectangle(2, 2, 5, 5)
	C, _ := NewRectangle(6, 6, 8, 8)

	if A.Equals(A) != true {
		t.Errorf("Equals test failed for rect: %v Equals: %v, Expected: %v, but got: %v", A, A, true, false)
	}

	if B.Equals(B) != true {
		t.Errorf("Equals test failed for rect: %v Equals: %v, Expected: %v, but got: %v", B, B, true, false)
	}

	if A.Equals(B) != false {
		t.Errorf("Equals test failed for rect: %v Equals: %v, Expected: %v, but got: %v", A, B, false, true)
	}

	if C.Equals(B) != false {
		t.Errorf("Equals test failed for rect: %v Equals: %v, Expected: %v, but got: %v", C, B, false, true)
	}
}

func TestExtendRectangle(t *testing.T) {
	// Define rectangles for test cases
	A, _ := NewRectangle(1, 1, 4, 4)
	B, _ := NewRectangle(2, 2, 5, 5)
	C, _ := NewRectangle(6, 6, 8, 8)
	D, _ := NewRectangle(3, 3, 7, 7)
	E, _ := NewRectangle(10, 10, 12, 12)
	F, _ := NewRectangle(5, 5, 20, 20)

	// Define test cases with clear expected results
	testCases := []struct {
		rect1, rect2 *Rectangle
		expected     *Rectangle
	}{
		{rect1: A, rect2: B, expected: &Rectangle{1, 1, 5, 5}},
		{rect1: C, rect2: D, expected: &Rectangle{3, 3, 8, 8}},
		{rect1: E, rect2: F, expected: &Rectangle{5, 5, 20, 20}},
	}

	for _, tc := range testCases {
		// Make a copy of rect1 to avoid in-place mutation affecting later tests
		rect1Copy := *tc.rect1

		// Extend rect1 by rect2
		rect1Copy.Extend(tc.rect2)

		// Check if the result matches the expected rectangle
		if !rect1Copy.Equals(tc.expected) {
			t.Errorf("Extend test failed for rect1: %v, rect2: %v. Expected: %v, got: %v", tc.rect1, tc.rect2, tc.expected, rect1Copy)
		}
	}
}
