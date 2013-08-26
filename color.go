package main

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func (c Color) Towards(c2 *Color, amount float64) Color {
	return Color{
		c.R*(1-amount) + c2.R*amount,
		c.G*(1-amount) + c2.G*amount,
		c.B*(1-amount) + c2.B*amount,
		c.A*(1-amount) + c2.A*amount}
}

var Black Color = Color{0, 0, 0, 1}
var Red Color = Color{1, 0, 0, 1}
var Green Color = Color{0, 1, 0, 1}
var Blue Color = Color{0, 0, 1, 1}
var White Color = Color{1, 1, 1, 1}
