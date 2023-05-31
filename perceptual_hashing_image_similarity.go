package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/corona10/goimagehash"
)

func main() {
	// Open the image file
	file1, err := os.Open("img/toped-watermark.png")
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file1.Close()

	// Decode the image file
	img1, _, err := image.Decode(file1)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	// Open the image file
	file2, err := os.Open("img/toped.png")
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file2.Close()

	// Decode the image file
	img2, _, err := image.Decode(file2)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	hash1, _ := goimagehash.ExtPerceptionHash(img1, 8, 8)
	hash2, _ := goimagehash.ExtPerceptionHash(img2, 8, 8)
	distance, _ := hash1.Distance(hash2)
	fmt.Printf("Distance between images: %v, hash1: %+v, hash2: %+v\n", distance, hash1.ToString(), hash2.ToString())

}
