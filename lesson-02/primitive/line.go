package primitive

// Line struct
type Line struct {
	x, y   int
	lx, ly int
	color  string
}

// CreateNewLine function
func CreateNewLine(x, y, lx, ly int, c string) *Line {
	return &Line{
		x:     x,
		y:     y,
		lx:    lx,
		ly:    ly,
		color: c,
	}
}

// Move function
func (l *Line) Move(dx, dy int) {
	l.x += dx
	l.y += dy
}
