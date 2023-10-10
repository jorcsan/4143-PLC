// Jorge Santos
// PLC-4143
// program 3
package main

import (
	"fmt"

	"github.com/jorcsan/mymodule" // Adjust the import path based on your actual module path
)

func main() {
	imageURL := "https://cdn.cloudflare.steamstatic.com/steam/apps/1817070/ss_dfba6f2477bfa42be69ddfdffbd421d3943d20bf.1920x1080.jpg?t=1695916105" // Replace with your actual image URL
	fileName := "downloaded_image.jpg"                                                                                                             // Specify the desired file name

	//will download a image file using its url
	err := mymodule.GetImage(imageURL, fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := "Spider-Man"
	outputPath := "output.png"

	//will print out colored text to a output file
	err = mymodule.PrintColor(text, outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image saved as", outputPath)

	err = mymodule.Pixels(outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	output2 := "grayspider.png"

	err = mymodule.Grayscale(fileName, output2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
