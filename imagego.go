package imagego

func main() {
	img, err := image.openImage("/home/caio_tavares/Documents/github/imageGo/assets/subaru.jpeg")
	if err != nil {
		return
	}

	pixels, err := pkg.ImageTensor(img)
	if err != nil {
		return
	}

	println(img.ColorModel())
}
