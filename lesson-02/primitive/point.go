package primitive

// Point struct
type Point struct {
	x, y  int
	color string
}

// CreateNewPoint функция создает новый объект типа Point
func CreateNewPoint(x, y int, c string) *Point {
	return &Point{
		x:     x,
		y:     y,
		color: c,
	}
}

// Move function
func (p *Point) Move(dx, dy int) {
	p.x += dx
	p.y += dy
}

// String function
func (p *Point) String() string {
	return "I'm a point"
}
