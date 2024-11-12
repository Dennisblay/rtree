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
	// if r.xMin > other.xMax || other.xMin > r.xMax {
	// 	return false
	// }
	//
	// // If rectangle is on top of top edge of other rectangle
	// if r.yMin > other.yMax || other.yMin > r.yMax {
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

func (r *Rectangle) Area() float64 {
	return (r.maxX - r.minX) * (r.maxY - r.minY)
}

func (r *Rectangle) Extend(other *Rectangle) {
	r.minX = math.Min(r.minX, other.minX)
	r.minY = math.Min(r.minY, other.minY)
	r.maxX = math.Max(r.maxX, other.maxX)
	r.maxY = math.Max(r.maxY, other.maxY)
}
