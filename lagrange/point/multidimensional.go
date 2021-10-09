package point

type MultidimensionalPoint struct {
	Point
	coordinates []uint8
}

// Returns Space Size
func (p *MultidimensionalPoint) Mes() int {
	return len(p.coordinates)
}
