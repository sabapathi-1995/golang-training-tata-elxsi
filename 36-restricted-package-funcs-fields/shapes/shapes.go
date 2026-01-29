package shapes

// Anything from a package stats with UpperCase is considerd as exportable, unrestricted

type Rect struct {
	L, B float32
	a, p float64
}

func newRect(l, b float32) *Rect {
	return &Rect{l, b, 0, 0}
}

func (r *Rect) Area() float64 {
	r.a = float64(r.L * r.B)
	return r.a
}

func (r *Rect) perimeter() float64 {
	r.p = float64(2 * (r.L + r.B))
	return r.p
}

type square struct {
	Side float32
}

func (r *square) Area() float64 {

	return float64(r.Side)
}

func (r *square) perimeter() float64 {
	return float64(r.Side) * 4
}

func NewSquare(s float32) *square {
	return &square{s}
}

func greet() {
	println("Welcome to shapes package")
}

func What() {
	println("The pacakge is for dealing with shapes like Rect and square")
}
