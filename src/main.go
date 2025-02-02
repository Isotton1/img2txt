package main

import (
	"os"
	"fmt"
	"log"
	"image"
	_ "image/png"
	_ "image/jpeg"
	_ "golang.org/x/image/webp"
	"strings"
)

func img2txt(img image.Image, width, height int) string {
	//https://paulbourke.net/dataformats/asciiart/
	const grayscale = " .:=+*%@"
	
	img_w := img.Bounds().Dx()
	img_h := img.Bounds().Dy()
	width = int(float64(80) * (float64(img_w)/float64(img_h))) 
	height = 35

	var sb strings.Builder
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			x := int(float64(w) / float64(width) * float64(img_w))
            y := int(float64(h) / float64(height) * float64(img_h))
			r, g, b, _ := img.At(x, y).RGBA()
			sb.WriteByte(grayscale[int(((r + g + b) / 24576))]); // /3 /65536 *8
		}
		sb.WriteByte('\n');
	}
	return sb.String()
}

func main() {
	argv := os.Args
	file, err := os.Open(argv[1])
	if (err != nil) {
		log.Fatal(err);
	}
	defer file.Close()
	
	img, _, err := image.Decode(file);
	if (err != nil) {
		log.Fatal(err);
	}

	fmt.Printf("%s", img2txt(img, 0, 0))
}
