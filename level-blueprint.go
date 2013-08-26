package main

import (
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type PerPlayerLevelBlueprint struct {
	Orientation Orientation3

	PaddleAreaCenter Point3
	PaddleAreaWidth  float64
	PaddleAreaHeight float64

	TargetAreaCenter Point3
}

type VertexBlueprint struct {
	Position Point3
	Color    Color
	Tu       float64
	Tv       float64
}

type FaceBlueprint struct {
	Vertices     []VertexBlueprint
	TextureIndex int
}

type LevelBlueprint struct {
	Players []PerPlayerLevelBlueprint
	Walls   []FaceBlueprint
}

func SetUpCamera(
	screenWidth int32,
	screenHeight int32,
	lb *LevelBlueprint,
	playerIndex int32) {
	pb := &lb.Players[playerIndex]

	// Calculate the distance between the near plane and the camera. We want
	// (PaddleAreaWidth / 2) / cameraDist = tan(CameraHorzFovDegrees / 2)
	t := math.Tan(math.Pi * Config.CameraHorzFovDegrees / 360)
	cameraDist := pb.PaddleAreaWidth / (2 * t)
	cameraPos := pb.PaddleAreaCenter.Minus(
		pb.Orientation.Forward.Times(cameraDist))

	// Set up the projection matrix.
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, gl.Sizei(screenWidth), gl.Sizei(screenHeight))
	gl.Frustum(
		gl.Double(-pb.PaddleAreaWidth/2), gl.Double(pb.PaddleAreaWidth/2),
		gl.Double(-pb.PaddleAreaHeight/2), gl.Double(pb.PaddleAreaHeight/2),
		gl.Double(cameraDist), gl.Double(cameraDist+Config.ViewDepth))

	// Set up the modelview matrix.
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	transformation := [16]gl.Double{
		gl.Double(pb.Orientation.Right.X),
		gl.Double(pb.Orientation.Up.X),
		gl.Double(-pb.Orientation.Forward.X),
		gl.Double(0.0),
		gl.Double(pb.Orientation.Right.Y),
		gl.Double(pb.Orientation.Up.Y),
		gl.Double(-pb.Orientation.Forward.Y),
		gl.Double(0.0),
		gl.Double(pb.Orientation.Right.Z),
		gl.Double(pb.Orientation.Up.Z),
		gl.Double(-pb.Orientation.Forward.Z),
		gl.Double(0.0),
		gl.Double(0.0), gl.Double(0.0), gl.Double(0.0), gl.Double(1.0)}
	gl.MultMatrixd(&transformation[0])
	gl.Translated(
		gl.Double(-cameraPos.X),
		gl.Double(-cameraPos.Y),
		gl.Double(-cameraPos.Z))
}

func MakeLevelVertexBuffer(lb *LevelBlueprint) *VertexBuffer {
	vb := VertexBuffer{}
	points := []Point3{}
	colors := []Color{}
	texCoords := [][2]float64{}

	for wi := range lb.Walls {
		indices := []gl.Int{}
		for vi := range lb.Walls[wi].Vertices {
			points = append(
				points, lb.Walls[wi].Vertices[vi].Position)
			colors = append(colors, lb.Walls[wi].Vertices[vi].Color)
			texCoords = append(
				texCoords,
				[2]float64{lb.Walls[wi].Vertices[vi].Tu, lb.Walls[wi].Vertices[vi].Tv})
			indices = append(indices, gl.Int(len(points)-1))
		}
		renderStep := VertexBufferRenderStep{gl.TRIANGLE_FAN, indices}
		vb.RenderSteps = append(vb.RenderSteps, renderStep)
	}

	vb.InitVertexData(points, colors, texCoords)
	return &vb
}
