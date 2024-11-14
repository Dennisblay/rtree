package rtree

import (
	"fmt"
	"math"
)

// Rectangle stores coordinates of a rectangle
type Rectangle struct {
	minX, minY, maxX, maxY float64
}

func (r *Rectangle) Overlaps(other *Rectangle) bool {
	// If rectangle is on right of the right edge of other rectangle
	// if r.minX > other.maxX || other.minX > r.maxX {
	// 	return false
	// }
	//
	// // If rectangle is on top of top edge of other rectangle
	// if r.minY > other.maxY || other.minY > r.maxY {
	// 	return false
	// }
	// return true

	// Putting them together
	return !(r.minX > other.maxX || other.minX > r.maxX || r.minY > other.maxY || other.minY > r.maxY)
}

func (r *Rectangle) Contains(other *Rectangle) bool {
	return r.minX <= other.minX && r.minY <= other.minY && r.maxX >= other.maxX && r.maxY >= other.maxY
}

func NewRectangle(coords ...float64) (*Rectangle, error) {
	if len(coords) > 4 {
		return nil, fmt.Errorf("NewRectangle requires exactly 4 coordinates, got %d", len(coords))
	}
	return &Rectangle{minX: coords[0], minY: coords[1], maxX: coords[2], maxY: coords[3]}, nil
}

// Centroid Computes the centroid of a Rectangle
func (r *Rectangle) Centroid() (float64, float64) {
	return r.minX + (r.maxX-r.minX)/2, r.minY + (r.maxY-r.minY)/2
}

// Distance computes euclidean distance between centers of two rectangles
func (r *Rectangle) Distance(other *Rectangle) float64 {
	centerX1, centerY1 := r.Centroid()
	centerX2, centerY2 := other.Centroid()
	return math.Sqrt((centerX2-centerX1)*(centerX2-centerX1) + (centerY2-centerY1)*(centerY2-centerY1))

}

// Area returns the Area of a Rectangle
func (r *Rectangle) Area() float64 {
	return (r.maxX - r.minX) * (r.maxY - r.minY)
}

// Extend extends the bounding box of a rectangle with another rectangle in place
func (r *Rectangle) Extend(other *Rectangle) {
	r.minX = math.Min(r.minX, other.minX)
	r.minY = math.Min(r.minY, other.minY)
	r.maxX = math.Max(r.maxX, other.maxX)
	r.maxY = math.Max(r.maxY, other.maxY)
}

func (r Rectangle) Union(other *Rectangle) *Rectangle {
	return &Rectangle{
		minX: math.Min(r.minX, other.minX),
		minY: math.Min(r.minY, other.minY),
		maxX: math.Max(r.maxX, other.maxX),
		maxY: math.Max(r.maxY, other.maxY),
	}
}

// Equals checks for equality with another rectangle
func (r *Rectangle) Equals(other *Rectangle) bool {
	return r.minX == other.minX && r.minY == other.minY && r.maxX == other.maxX && r.maxY == other.maxY
}
