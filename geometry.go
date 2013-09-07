package main

import (
	"math"
)

type Point3 struct {
	X float64
	Y float64
	Z float64
}

func (p Point3) Plus(p2 Point3) Point3 {
	return Point3{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

func (p Point3) Minus(p2 Point3) Point3 {
	return Point3{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

func (p Point3) Times(scale float64) Point3 {
	return Point3{p.X * scale, p.Y * scale, p.Z * scale}
}

func (p Point3) Dot(p2 Point3) float64 {
	return p.X*p2.X + p.Y*p2.Y + p.Z*p2.Z
}

func (p Point3) Cross(p2 Point3) Point3 {
	return Point3{
		p.Y*p2.Z - p.Z*p2.Y,
		p.Z*p2.X - p.X*p2.Z,
		p.X*p2.Y - p.Y*p2.X}
}

func (p Point3) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

type Orientation3 struct {
	Right   Point3
	Up      Point3
	Forward Point3
}

type Polygon3 struct {
	Points []Point3
}
