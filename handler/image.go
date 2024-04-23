package handler

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func OpenImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	// defer f.Close()
	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}
	if format != "jpeg" {
		fmt.Println("image format is not jpeg")
		return nil, errors.New("")
	}

	return img, nil
}

func ImageToTensor(img image.Image) ([][]color.Color, error) {
	// Converting image into tensor
	size := img.Bounds().Size()
	var pixels [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels, nil
}

func TensorToImage(pixels [][]color.Color) (image.Image, error) {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}

	}

	return nImg, nil
}

func SaveImage(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return err
	}

	return nil
}
