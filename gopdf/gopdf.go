package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	// Create a new PDF object
	pdf := gopdf.GoPdf{}
	// Set page size to A4 (595.28 x 841.89 points)
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Add a new page
	pdf.AddPage()

	// Add a font (TTF required)
	err := pdf.AddTTFFont("arial", "C:/Windows/Fonts/arial.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	// Set font
	err = pdf.SetFont("arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// Write text
	pdf.SetX(50)
	pdf.SetY(100)
	pdf.Cell(nil, "Hello, this is a PDF created with Go!")

	// Save to file
	err = pdf.WritePdf("example.pdf")
	if err != nil {
		log.Print(err.Error())
	}
}
