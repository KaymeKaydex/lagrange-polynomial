package point

import (
	"fmt"
	"sync/atomic"
)

type Point2D struct {
	Point
	x, y int64
}

// Mes Returns Space Size
func (p *Point2D) Mes() int {
	return Mes2D
}

func (p *Point2D) String() string {
	return fmt.Sprintf("x:%d y:%d", p.x, p.y)
}

func (p *Point2D) GetX() int64 {
	return p.x
}
func (p *Point2D) GetY() int64 {
	return p.x
}

func (p *Point2D) SetX(value int64) {
	atomic.StoreInt64(&p.x, value)
}

func (p *Point2D) SetY(value int64) {
	atomic.StoreInt64(&p.y, value)
}
