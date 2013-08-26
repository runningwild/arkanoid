package main

import (
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type Ball struct {
	Position Point3
	Velocity Point3

	vertexBuffer VertexBuffer
}

func MakeBall() Ball {
	b := Ball{}
	b.vertexBuffer = makeBallVertexBuffer()
	b.Position = Point3{0, 0, 1}
	b.Velocity = Point3{0, 0, 0.1}
	return b
}

func makeBallVertexBuffer() VertexBuffer {
	vb := VertexBuffer{}
	points := []Point3{}
	colors := []Color{}
	texCoords := [][2]float64{}
	indices := []gl.Int{}

	for i := 0; i < 20; i++ {
		theta := (2 * math.Pi) * (float64(i) / 20.0)
		points = append(points, Point3{
			Config.BallRadius * math.Cos(theta),
			Config.BallRadius * math.Sin(theta), 0})
		colors = append(colors, Black)
		texCoords = append(texCoords, [2]float64{0, 0})
		indices = append(indices, gl.Int(i))
	}

	vb.InitVertexData(points, colors, texCoords)
	renderStep := VertexBufferRenderStep{gl.LINE_LOOP, indices}
	vb.RenderSteps = []VertexBufferRenderStep{renderStep}
	return vb
}

func (b *Ball) Render() {
	b.vertexBuffer.Render(b.Position)
}
