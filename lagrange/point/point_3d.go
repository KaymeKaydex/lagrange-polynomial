package point

type Point3D struct {
	Point
	x, y, z uint8
}

// Returns Space Size
func (p *Point3D) Mes() int {
	return Mes3D
}
