package rect

type Rect struct {
	L, B float32
}

func New(l, b float32) *Rect {
	return &Rect{l, b}
	// return &Rect{L:l,B:b}
}

func (r *Rect) Area() float64 {
	return float64(r.L * r.B)
}

func (r *Rect) Perimeter() float64 {
	return 2 * float64(r.L+r.B)
}

func (*Rect) What() string { // r is not being used , can just create a receiver but not a variable
	return "Rect"
}
