package calculator

type Calculator struct {
	NumberOne int
	NumberTwo int
}

func (c Calculator) Add() int {
	return c.NumberOne + c.NumberTwo
}

func (c Calculator) Multiply() int {
	return c.NumberOne * c.NumberTwo
}

func (c Calculator) Divide() int {
	if c.NumberTwo == 0 || c.NumberOne%c.NumberTwo > 0 {
		return 0
	}
	return c.NumberOne / c.NumberTwo
}
