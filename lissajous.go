package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"
	"os"
	"image/gif"
	"math"
	"io"
)

var palette = []color.Color{color.White, color.RGBA{0, 255, 0, 255}, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer)  {
	const (
		cycle = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount:nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i % (len(palette) - 1)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}