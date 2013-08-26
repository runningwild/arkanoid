package main

type ConfigType struct {
	// Sizes.
	WindowWidth  int
	WindowHeight int

	ViewDepth            float64
	CameraHorzFovDegrees float64

	BlockWidth  float64
	BlockHeight float64
	BlockDepth  float64

	BallRadius float64

	TargetAreaXBlocks int
	TargetAreaYBlocks int
	TargetAreaZBlocks int

	BasePaddleWidth  float64
	BasePaddleHeight float64
	PaddleDepth      float64

	// Colors.
	PaddleBackColor      Color
	PaddleFrontColor     Color
	PaddleBackEdgeColor  Color
	PaddleFrontEdgeColor Color

	// Times (always in seconds).
	FrameTime float64
}

var Config ConfigType
