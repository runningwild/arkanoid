package main

import (
	gl "github.com/chsc/gogl/gl21"
)

type VertexBufferRenderStep struct {
	Mode    gl.Enum
	Indices []gl.Int
}

type VertexBuffer struct {
	VertexData   []gl.Double
	ColorData    []gl.Double
	TexCoordData []gl.Double
	RenderSteps  []VertexBufferRenderStep
}

func (vb *VertexBuffer) InitVertexData(
	points []Point3,
	colors []Color,
	texCoords [][2]float64) {
	vb.VertexData = make([]gl.Double, len(points)*3)
	for i := range points {
		vb.VertexData[3*i+0] = gl.Double(points[i].X)
		vb.VertexData[3*i+1] = gl.Double(points[i].Y)
		vb.VertexData[3*i+2] = gl.Double(points[i].Z)
	}

	vb.ColorData = make([]gl.Double, len(colors)*4)
	for i := range colors {
		vb.ColorData[4*i+0] = gl.Double(colors[i].R)
		vb.ColorData[4*i+1] = gl.Double(colors[i].G)
		vb.ColorData[4*i+2] = gl.Double(colors[i].B)
		vb.ColorData[4*i+3] = gl.Double(colors[i].A)
	}

	vb.TexCoordData = make([]gl.Double, len(texCoords)*2)
	for i := range texCoords {
		vb.TexCoordData[2*i+0] = gl.Double(texCoords[i][0])
		vb.TexCoordData[2*i+1] = gl.Double(texCoords[i][1])
	}
}

func (vb *VertexBuffer) Render(position Point3) {
	gl.PushMatrix()
	gl.Translated(
		gl.Double(position.X),
		gl.Double(position.Y),
		gl.Double(position.Z))

	gl.VertexPointer(
		gl.Int(3),
		gl.DOUBLE,
		0,
		gl.Pointer(&vb.VertexData[0]))
	gl.ColorPointer(
		gl.Int(4),
		gl.DOUBLE,
		0,
		gl.Pointer(&vb.ColorData[0]))
	gl.TexCoordPointer(
		gl.Int(2),
		gl.DOUBLE,
		0,
		gl.Pointer(&vb.TexCoordData[0]))

	for i := range vb.RenderSteps {
		rs := &vb.RenderSteps[i]
		gl.DrawElements(
			rs.Mode,
			gl.Sizei(len(rs.Indices)),
			gl.UNSIGNED_INT,
			gl.Pointer(&rs.Indices[0]))
	}
	gl.PopMatrix()
}
