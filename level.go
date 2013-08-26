package main

import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/runningwild/cgf"
	"github.com/runningwild/glop/system"
	"math"
)

type Level struct {
	LevelBlueprint    *LevelBlueprint
	levelVertexBuffer *VertexBuffer

	Paddle Paddle
	Ball   Ball
}

func MakeLevel(lb *LevelBlueprint) *Level {
	return &Level{
		lb,
		MakeLevelVertexBuffer(lb),
		MakePaddle(),
		MakeBall()}
}

func (l *Level) Think() {
	l.Ball.Position = l.Ball.Position.Plus(l.Ball.Velocity)
	l.Ball.Velocity.Y -= 0.002

	if l.Ball.Position.Y < -3 {
		l.Ball.Position.Y = -6 - l.Ball.Position.Y
		l.Ball.Velocity.Y = -l.Ball.Velocity.Y
	}

	if l.Ball.Position.Z < 0 {
		l.Ball.Position.Z = -l.Ball.Position.Z
		l.Ball.Velocity.Z = -l.Ball.Velocity.Z
	}

	if l.Ball.Position.Z > 10 {
		l.Ball.Position.Z = 20 - l.Ball.Position.Z
		l.Ball.Velocity.Z = -l.Ball.Velocity.Z
	}
}

func (l *Level) LocalThink(sys system.System, engine *cgf.Engine) {
	// Convert the mouse screen coordinates to paddleWidth, paddleHeight range.
	pixelX, pixelY := sys.GetCursorPos()
	pb := &l.LevelBlueprint.Players[0]
	transformedX := float64(pixelX-Config.WindowWidth/2) *
		(pb.PaddleAreaWidth / float64(Config.WindowWidth))
	transformedY := -float64(pixelY-Config.WindowHeight/2) *
		(pb.PaddleAreaHeight / float64(Config.WindowHeight))

	// Clip this to a valid range.
	p := &l.Paddle
	x := math.Max(-pb.PaddleAreaWidth/2+p.Width/2,
		math.Min(transformedX, pb.PaddleAreaWidth/2-p.Width/2))
	y := math.Max(-pb.PaddleAreaHeight/2+p.Height/2,
		math.Min(transformedY, pb.PaddleAreaHeight/2-p.Height/2))

	// Create the event.
	engine.ApplyEvent(&MovePlayerEvent{x, y})
}

func (l *Level) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	SetUpCamera(800, 600, l.LevelBlueprint, 0)
	l.levelVertexBuffer.Render(Point3{0, 0, 0})
	l.Paddle.Render()
	l.Ball.Render()
}
