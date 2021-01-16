package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	config := pixelgl.WindowConfig{
		Title:  "teste",
		Bounds: pixel.R(0, 0, 768, 720), //256, 240
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(config)
	if err != nil {
		panic(err)
	}

	//imd := imdraw.New(nil)

	colors := [...]color.RGBA{
		colornames.Red,
		colornames.Blue,
		colornames.Green,
		colornames.Aquamarine,
		colornames.Bisque,
		colornames.White,
		colornames.Black,
		colornames.Purple,
		colornames.Chartreuse,
		colornames.Cyan,
		colornames.Blanchedalmond,
		colornames.Darkgoldenrod,
	}

	rand.Seed(time.Now().UnixNano())

	imageChannel := make(chan *imdraw.IMDraw, 9)
	defer close(imageChannel)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				go fillScreen(256*j, 720-(240*i), colors, imageChannel)
			}
		}

		drawSlice := make([]*imdraw.IMDraw, 0)

		for i := 0; i < 9; i++ {
			drawSlice = append(drawSlice, <-imageChannel)
		}

		for _, image := range drawSlice {
			image.Draw(win)
		}
		win.Update()
	}
}

func fillScreen(xExt, yExt int, colors [12]color.RGBA, onii chan *imdraw.IMDraw) {

	imd := imdraw.New(nil)

	for y := yExt; y >= yExt-240; y -= 3 {
		for x := xExt; x < xExt+256; x += 3 {
			imd.Color = colors[rand.Int()%12]
			imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+3), float64(y-3)))
			imd.Rectangle(0)
		}
	}

	onii <- imd

}
