package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("testdata/org.jpg")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("Could not decode image: %s", err)
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Could not create file")
	}

	defer outputFile.Close()

	asciiBrightnessChars := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,^."

	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 40 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 10 {
			r, g, b, _ := m.At(x, y).RGBA()

			combined := ((0.2126 * float64(r)) + (0.7152 * float64(g)) + (0.0722 * float64(b))) / 257

			resultIndex := int((combined / 255) * float64(len(asciiBrightnessChars)-1))
			outputFile.WriteString(string(asciiBrightnessChars[resultIndex]))

		}
		outputFile.WriteString("\n")
	}

	fmt.Println("Saved to output.txt")
}
