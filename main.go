package main

import (
	"encoding/json"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/runningwild/cgf"
	"github.com/runningwild/glop/gin"
	"github.com/runningwild/glop/gos"
	"github.com/runningwild/glop/render"
	"github.com/runningwild/glop/system"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

func loadJson(filename string, v interface{}) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
}

func init() {
	// Required for Darwin.
	runtime.LockOSThread()

	loadJson("data/config.json", &Config)
}

func initWindow(sys system.System, width int, height int) {
	sys.CreateWindow(10, 10, width, height)
	sys.EnableVSync(false)

	err := gl.Init()
	if err != nil {
		panic(err)
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.COLOR_ARRAY)
	gl.Enable(gl.CULL_FACE)
	gl.FrontFace(gl.CW)

	gl.ClearColor(0, 0, 0, 1)
}

func main() {
	runtime.GOMAXPROCS(2)
	sys := system.Make(gos.GetSystemInterface())
	sys.Startup()

	game := Game{}
	var lb LevelBlueprint
	loadJson("data/1p_basic_level.json", &lb)
	if len(lb.Players) == 0 || len(lb.Walls) == 0 {
		panic(fmt.Sprintf("Invalid level config: %d players and %d walls.",
			len(lb.Players), len(lb.Walls)))
	}
	engine, _ := cgf.NewLocalEngine(&game, int(Config.FrameTime*1000), nil)
	engine.ApplyEvent(&NewLevelEvent{&lb})

	render.Init()
	render.Queue(func() {
		initWindow(sys, Config.WindowWidth, Config.WindowHeight)
	})
	render.Purge()

	ticker := time.Tick(time.Millisecond * time.Duration(Config.FrameTime*1000))
	for true {
		<-ticker
		LocalThink(sys, engine, &game)
		if gin.In().GetKey(gin.AnyEscape).FramePressCount() > 0 {
			break
		}
	}
}
