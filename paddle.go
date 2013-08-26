package main

import (
	gl "github.com/chsc/gogl/gl21"
)

type Paddle struct {
	Width        float64
	Height       float64
	X            float64
	Y            float64
	vertexBuffer *VertexBuffer
}

func MakePaddle() Paddle {
	p := Paddle{}
	p.Width = Config.BasePaddleWidth
	p.Height = Config.BasePaddleHeight
	p.X = 0
	p.Y = 0
	p.vertexBuffer = makePaddleVertexBuffer()
	return p
}

func makePaddleVertexBuffer() *VertexBuffer {
	points := []Point3{
		// Paddle face points.
		Point3{-Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{-Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{-Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2, 0},
		Point3{Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2, 0},
		Point3{Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2, 0},
		Point3{-Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2, 0},
		// Paddle edge points.
		Point3{-Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{-Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2,
			-Config.PaddleDepth},
		Point3{-Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2, 0},
		Point3{Config.BasePaddleWidth / 2, Config.BasePaddleHeight / 2, 0},
		Point3{Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2, 0},
		Point3{-Config.BasePaddleWidth / 2, -Config.BasePaddleHeight / 2, 0}}

	colors := []Color{
		Config.PaddleBackColor,
		Config.PaddleBackColor,
		Config.PaddleBackColor,
		Config.PaddleBackColor,
		Config.PaddleFrontColor,
		Config.PaddleFrontColor,
		Config.PaddleFrontColor,
		Config.PaddleFrontColor,
		Config.PaddleBackEdgeColor,
		Config.PaddleBackEdgeColor,
		Config.PaddleBackEdgeColor,
		Config.PaddleBackEdgeColor,
		Config.PaddleFrontEdgeColor,
		Config.PaddleFrontEdgeColor,
		Config.PaddleFrontEdgeColor,
		Config.PaddleFrontEdgeColor}

	texCoords := [][2]float64{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	vb := VertexBuffer{}
	vb.InitVertexData(points, colors, texCoords)
	vb.RenderSteps = []VertexBufferRenderStep{
		VertexBufferRenderStep{gl.QUADS, []gl.Int{
			0, 1, 2, 3,
			4, 0, 3, 7,
			1, 5, 6, 2,
			4, 5, 1, 0,
			3, 2, 6, 7,
			5, 4, 7, 6,
		}},
		VertexBufferRenderStep{gl.LINES, []gl.Int{
			8, 9, 11, 10, 12, 13, 15, 14,
			8, 11, 9, 10, 12, 15, 13, 14,
			8, 12, 9, 13, 11, 15, 10, 14,
		}},
	}
	return &vb
}

func (p *Paddle) Render() {
	p.vertexBuffer.Render(Point3{p.X, p.Y, Config.PaddleDepth})
}
