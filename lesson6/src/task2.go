package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	DrawLine(rectImg, 100, 50, 100, red, true)
	DrawLine(rectImg, 50, 100, 100, red, false)

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}

func DrawLine(img *image.RGBA, startX, startY, length int, color color.RGBA, isVertival bool) {
	s := startX
	l := startY
	if isVertival {
		l, s = s, l
	}
	e := s + length
	for ; s <= e; s++ {
		if isVertival {
			img.Set(l, s, color)
		} else {
			img.Set(s, l, color)
		}
	}
}
