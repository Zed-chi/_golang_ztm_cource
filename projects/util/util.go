package util
import "image"
import "image/color"

func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors:= make(map[color.Color]struct{})

	var empty struct{}

	for y := 0; y<img.Bounds().Dy(); y++{
		for x:=;x<img.Bounds().Dx();x++{
			colors[img.At(x,y)] = empty

		}
	}
	return colors
}