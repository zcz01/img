package img

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"

	"os"
	"path/filepath"
	"strings"

	bmp "golang.org/x/image/bmp"
)

/// gray
func ToGray(AImage image.Image) *image.Gray {
	bounds := AImage.Bounds()
	reGray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			reGray.Set(x, y, AImage.At(x, y))
		}
	}
	return reGray
}

/// block
func ToBlock(AGray *image.Gray) *image.Gray {
	bounds := AGray.Bounds()
	reGray := image.NewGray(bounds)
	var oldPixel, newPixel color.Gray
	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			oldPixel = AGray.GrayAt(i, j)
			if oldPixel.Y < 143 {
				newPixel = color.Gray{0}
			} else {
				newPixel = color.Gray{255}
			}
			reGray.SetGray(i, j, newPixel)
		}
	}
	return reGray
}

/// Halftone
func ToHalftone(AGray *image.Gray) *image.Gray {
	bounds := AGray.Bounds()
	reGray := image.NewGray(bounds)
	width := bounds.Dx()
	height := bounds.Dy()
	var oldPixel, newPixel color.Gray
	var t int
	copy(reGray.Pix, AGray.Pix)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			oldPixel = reGray.GrayAt(i, j)
			if oldPixel.Y < 158 {
				newPixel = color.Gray{0}
			} else {
				newPixel = color.Gray{255}
			}
			reGray.SetGray(i, j, newPixel)
			t = (int(oldPixel.Y) - int(newPixel.Y)) / 16
			reGray.SetGray(i+1, j, gray2(reGray.GrayAt(i+1, j).Y, 7*t))
			reGray.SetGray(i-1, j+1, gray2(reGray.GrayAt(i-1, j+1).Y, 3*t))
			reGray.SetGray(i, j+1, gray2(reGray.GrayAt(i, j+1).Y, 5*t))
			reGray.SetGray(i+1, j+1, gray2(reGray.GrayAt(i+1, j+1).Y, t))
		}
	}
	return reGray
}

func gray2(y uint8, t int) color.Gray {
	x := int(y) + t
	switch {
	case x < 1:
		return color.Gray{0}
	case x > 254:
		return color.Gray{255}
	}
	return color.Gray{uint8(x)}
}

/// open
func Open(image_file string) (Image image.Image, err error) {
	f, err := os.Open(image_file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	ext := strings.ToLower(filepath.Ext(image_file))
	if "jpg" == ext || "jpeg" == ext {
		Image, err = jpeg.Decode(f)
	} else if "png" == ext {
		Image, err = png.Decode(f)
	} else if "bmp" == ext {
		Image, err = bmp.Decode(f)
	} else if "gif" == ext {
		Image, err = gif.Decode(f)
	} else {
		Image, _, err = image.Decode(f)
	}
	return
}

func SavePNG(AImage image.Image, save_file string) error {
	f, err := os.Create(save_file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, AImage)
	if err != nil {
		return err
	}
	return nil
}

func SaveJPG(AImage image.Image, save_file string) error {
	f, err := os.Create(save_file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, AImage, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}

func SaveBMP(AImage image.Image, save_file string) error {
	f, err := os.Create(save_file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = bmp.Encode(f, AImage)
	if err != nil {
		return err
	}
	return nil
}
