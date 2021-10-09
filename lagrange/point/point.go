package point

type Point interface {
	Mes() int
	String() string
	Move(...uint8) error
}
