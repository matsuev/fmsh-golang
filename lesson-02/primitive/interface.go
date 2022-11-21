package primitive

import "errors"

// Primitive interface
type Primitive interface {
	Move(dx, dy int)
}

func CreateNewPrimitive(pType string, pColor string, x, y int, args ...int) (Primitive, error) {
	switch pType {
	case "point":
		p := CreateNewPoint(x, y, pColor)
		return p, nil
	case "line":
		if len(args) < 2 {
			return nil, errors.New("not enought arguments")
		}
		return &Line{
			x:     x,
			y:     y,
			lx:    args[0],
			ly:    args[1],
			color: pColor,
		}, nil
	}
	return nil, nil
}
