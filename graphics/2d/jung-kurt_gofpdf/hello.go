package main

import (
	"github.com/jung-kurt/gofpdf"
	"image"
	"log"
	"os"
)

func main() {
	// Create a new PDF with Letter size and landscape orientation
	pdf := gofpdf.New("L", "pt", "Letter", "")

	// Define points per inch (adjust this based on your printer scaling if needed)
	pointsPerInch := 77.0

	// Define rectangle dimensions
	rectWidth := 2 * pointsPerInch  // 2 inches wide
	rectHeight := 6 * pointsPerInch // 6 inches tall
	margin := 0.0                   // No margin between rectangles

	// Load the image to get its aspect ratio
	imagePath := "./graphics/2d/jung-kurt_gofpdf/sample_image.jpg"
	imgFile, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("Failed to open image file: %v", err)
	}
	defer imgFile.Close()

	img, _, err := image.DecodeConfig(imgFile)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	// Calculate the image aspect ratio
	imageAspectRatio := float64(img.Width) / float64(img.Height)

	// Determine the image dimensions to maintain aspect ratio within the rectangle
	var imageWidth, imageHeight float64
	if imageAspectRatio >= 1 {
		// Wider image
		imageWidth = rectWidth - 20 // Leave some padding
		imageHeight = imageWidth / imageAspectRatio
		if imageHeight > rectHeight-40 {
			imageHeight = rectHeight - 40
			imageWidth = imageHeight * imageAspectRatio
		}
	} else {
		// Taller image
		imageHeight = rectHeight - 40
		imageWidth = imageHeight * imageAspectRatio
		if imageWidth > rectWidth-20 {
			imageWidth = rectWidth - 20
			imageHeight = imageWidth / imageAspectRatio
		}
	}

	// Add a page
	pdf.AddPage()

	// Set background color to white
	pdf.SetFillColor(255, 255, 255)

	// Loop to create 4 repeating rectangles with image and text
	for i := 0; i < 4; i++ {
		// Calculate the X position for each rectangle
		x := float64(i) * rectWidth

		// Draw the rectangle border
		pdf.Rect(x, margin, rectWidth, rectHeight, "D")

		// Add the image inside the rectangle, maintaining aspect ratio
		imageX := x + (rectWidth-imageWidth)/2 // Center the image horizontally
		imageY := margin + 20                  // Set Y position (20 points from the top)
		opt := gofpdf.ImageOptions{
			ImageType:             "JPG",
			ReadDpi:               false,
			AllowNegativePosition: false,
		}
		pdf.ImageOptions(imagePath, imageX, imageY, imageWidth, imageHeight, false, opt, 0, "")

		// Add some text below the image
		pdf.SetFont("Arial", "", 12)
		textX := x + 10 // Indent text slightly from the left
		textY := imageY + imageHeight + 20
		pdf.Text(textX, textY, "Hello, World!")
	}

	// Output the PDF to a file
	err = pdf.OutputFileAndClose("./graphics/2d/jung-kurt_gofpdf/output.pdf")
	if err != nil {
		log.Fatalf("Error saving PDF: %v", err)
	}

	log.Println("PDF created successfully!")
}
