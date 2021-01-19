package img

import (
	"image"
	"math"

	"github.com/nfnt/resize"
)

/// Resize
func Resize(AImage image.Image, NewWidth, NewHeight float64) image.Image {
	b := AImage.Bounds()
	width := float64(b.Max.X)
	height := float64(b.Max.Y)
	ratio := math.Min(NewWidth/width, NewHeight/height)
	return resize.Resize(uint(width*ratio), uint(height*ratio), AImage, resize.Lanczos3)
}

 