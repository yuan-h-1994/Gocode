package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("/home/vagrant/share/up/GOcode/1.jpg")
	assert(err, "Invalid image file path ")
	defer file.Close()

	img, _, err := image.Decode(file)
	assert(err, "Failed to convert file to image.")

	out, err := os.Create("/home/vagrant/share/up/GOcode/png/1.png")
	assert(err, "Failed to create output path.")
	defer out.Close()

	png.Encode(out, img)
}

func assert(err error, msg string) {
	if err != nil {
		panic(err.Error() + ":" + msg)
	}
}
