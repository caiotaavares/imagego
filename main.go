package main

import (
	"fmt"

	"github.com/caiotaavares/imagego/handler"
)

func main() {
	imgPath := "assets/subaru.jpeg"
	img, err := handler.OpenImage(imgPath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}

	pixels, err := handler.ImageToTensor(img)
	if err != nil {
		fmt.Println("Erro no imageTensor:", err)
		return
	}

	newImg, err := handler.TensorToImage(pixels)
	handler.SaveImage(newImg, "assets/edited_subaru.jpeg")

	fmt.Println("Image opened successfully")
	fmt.Println(img.Bounds().Size())
}
