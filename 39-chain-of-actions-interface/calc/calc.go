package calc

func New(d int) *Calc {
	return &Calc{d}
}

type Calc struct {
	Data int
}

func (c *Calc) Add(d int) *Calc {
	c.Data += d
	return c
}

func (c *Calc) Sub(d int) *Calc {
	c.Data -= d
	return c

}

func (c *Calc) Mul(d int) *Calc {
	c.Data *= d
	return c

}

func (c *Calc) Div(d int) *Calc {
	c.Data /= d
	return c

}

func (c *Calc) Get() int {
	return c.Data
}

// Task is
// remove int and keep any . make sure that on all number types , all these operations shoul be implemented
