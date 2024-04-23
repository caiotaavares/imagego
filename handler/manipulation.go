package handler

import "image/color"

func UpsideDown(pixels [][]color.Color) ([][]color.Color, error) {
	height := len(pixels)
	width := len(pixels[0])
	result := make([][]color.Color, height)

	for i := 0; i < height; i++ {
		result[i] = make([]color.Color, width)
		for j := 0; j < width; j++ {
			result[i][j] = pixels[height-i-1][width-j-1]
		}
	}

	return result, nil
}
