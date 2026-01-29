package icalc

type ICalc interface {
	Add(d int) ICalc
	Sub(d int) ICalc
	Mul(d int) ICalc
	Div(d int) ICalc
	Get() int
}

func New(d int) *Calc {
	return &Calc{d}
}

type Calc struct {
	Data int
}

func (c *Calc) Add(d int) ICalc {
	c.Data += d
	return c
}

func (c *Calc) Sub(d int) ICalc {
	c.Data -= d
	return c

}

func (c *Calc) Mul(d int) ICalc {
	c.Data *= d
	return c

}

func (c *Calc) Div(d int) ICalc {
	c.Data /= d
	return c

}

func (c *Calc) Get() int {
	return c.Data
}

// Task is
// remove int and keep any . make sure that on all number types , all these operations shoul be implemented
