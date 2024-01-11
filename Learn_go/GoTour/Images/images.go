package main

import (
	"image"
	"image/color"
)

type Image struct {
    Width, Height int
    color uint8
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
    r := uint8(x + y)
    g := uint8(x ^ y)
    b := uint8(x * y)
    return color.RGBA{r, g, b, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
