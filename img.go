package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func IMAGE() {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			img.Set(x, y, color.Black)
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
