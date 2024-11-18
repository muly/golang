package main

import (
	"github.com/fogleman/gg"
	"log"
)

func main() {
	// Define DPI and rectangle dimensions
	const dpi = 340
	const rectWidth = 2 * dpi    // 2 inches wide
	const rectHeight = 6 * dpi   // 6 inches long
	const repeatCount = 4        // Number of rectangles

	// Calculate the total canvas width (side by side, no gaps)
	canvasWidth := repeatCount * rectWidth
	canvasHeight := rectHeight

	// Create a new drawing context
	dc := gg.NewContext(int(canvasWidth), int(canvasHeight))
	dc.SetRGB(1, 1, 1) // Set background color to white
	dc.Clear()

	// Load the default system font with a specific size
	fontPath := "/System/Library/Fonts/Supplemental/Arial.ttf" // Use Arial font on macOS
	if err := dc.LoadFontFace(fontPath, 36); err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	// Set the line width for borders
	dc.SetLineWidth(4)

	// Loop to draw 4 rectangles side by side
	for i := 0; i < repeatCount; i++ {
		x := float64(i) * rectWidth
		y := 0.0

		// Draw the rectangle with a border
		dc.SetRGB(1, 1, 1) // White fill color
		dc.DrawRectangle(x, y, rectWidth, rectHeight)
		dc.FillPreserve()  // Fill the rectangle and preserve the path for the border

		// Draw the border
		dc.SetRGB(0, 0, 0) // Black border color
		dc.Stroke()

		// Draw "Hello, World!" text
		dc.SetRGB(0, 0, 0) // Black text color
		dc.DrawStringAnchored("Hello, World!", x+rectWidth/2, y+rectHeight/2, 0.5, 0.5)
	}

	// Save the image as PNG
	outputFile := "output.png"
	if err := dc.SavePNG(outputFile); err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

	log.Printf("Image generated successfully: %s", outputFile)
}
