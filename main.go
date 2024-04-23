package main

import (
	"fmt"

	"github.com/caiotaavares/imagego/handler"
)

func main() {
	imgPath := "/home/caio_tavares/Documents/github/imagego/assets/subaru.jpeg"
	img, err := handler.OpenImage(imgPath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}

	pixels, err := handler.ImageTensor(img)
	if err != nil {
		fmt.Println("Erro no imageTensor:", err)
		return
	}

	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels); j++ {
		}
	}

	fmt.Println("Image opened successfully")
	fmt.Println(img.Bounds().Size())
}
