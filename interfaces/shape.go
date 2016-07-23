package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Hybrid struct {
	square Square
	rectangle Rectangle
}

type Rectangle struct {
	length, breadth int
}

func (s *Square) Area() int {
	return s.side * s.side
}

func (r *Rectangle) Area() int {
	return r.length * r.breadth
}

func (h *Hybrid) Area() int {
	return h.square.Area() + h.rectangle.Area()
}
