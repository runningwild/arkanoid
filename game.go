package main

import (
	"github.com/runningwild/cgf"
	"github.com/runningwild/glop/render"
	"github.com/runningwild/glop/system"
)

type Game struct {
	Level *Level
}

func (g *Game) Think() {
	if g.Level != nil {
		g.Level.Think()
	}
}

func (g *Game) Render() {
	g.Level.Render()
}

func LocalThink(sys system.System, engine *cgf.Engine, game *Game) {
	sys.Think()
	engine.Pause()

	// We might have no level if the event has not gone through yet.
	if game.Level != nil {
		render.Queue(func() {
			game.Render()
			sys.SwapBuffers()
		})
		game.Level.LocalThink(sys, engine)
	}

	render.Purge()
	engine.Unpause()
}
