package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("testdata/chuck.jpg")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Could not decode image: %s", err)
	}

	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 20 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 20 {
			// get all colors from image
			r, g, b, a := m.At(x, y).RGBA()

			combined := (0.2126 * float64(r)) + (0.7152 * float64(g)) + (0.0722*float64(b))/257
			fmt.Println(combined)
			fmt.Println(a)

		}
	}
}
