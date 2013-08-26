package main

import (
	"encoding/gob"
)

type NewLevelEvent struct {
	Lb *LevelBlueprint
}

func (nle *NewLevelEvent) Apply(_g interface{}) {
	g := _g.(*Game)
	g.Level = MakeLevel(nle.Lb)
}

type MovePlayerEvent struct {
	X float64
	Y float64
}

func (mpe *MovePlayerEvent) Apply(_g interface{}) {
	g := _g.(*Game)
	if g.Level != nil {
		g.Level.Paddle.X = mpe.X
		g.Level.Paddle.Y = mpe.Y
	}
}

func init() {
	gob.Register(NewLevelEvent{})
	gob.Register(MovePlayerEvent{})
}
