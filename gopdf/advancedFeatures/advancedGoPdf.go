package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

var (
	goPdfFolder      = "gopdfExample/"          // Base folder for all examples
	pdfCreation      = "pdfCreation/"           // Folder for PDF creation examples
	textHandling     = "textHandling/"          // Folder for text handling examples
	advancedFeatures = "advancedGopdfFeatures/" // Folder for advanced features examples
	//pdfPageSetup = "example.pdf"
)

func main() {
	// Create a new PDF document
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Add font with proper fallback
	fontName := "arial"
	err := pdf.AddTTFFont(fontName, "./fonts/arial.ttf")
	if err != nil {
		// Try system font on Windows
		err = pdf.AddTTFFont(fontName, "C:/Windows/Fonts/arial.ttf")
		if err != nil {
			// Use built-in PostScript font as final fallback
			fontName = "Times-Roman"
		}
	}

	// Example 1: Adding Images
	addImagesExample(&pdf)

	// Example 2: Drawing Shapes and Lines
	drawShapesExample(&pdf)

	// Example 3: Creating Tables and Grids
	createTableExample(&pdf)

	// Example 4: Headers and Footers
	addHeaderFooterExample(&pdf)

	// Example 5: Page Numbering
	addPageNumberingExample(&pdf)
}

// Example 1: Adding Images (PNG, JPEG)
func addImagesExample(pdf *gopdf.GoPdf) {
	pdf.AddPage()

	// Set font for the title
	pdf.SetFont("arial", "", 16)
	pdf.Cell(nil, "Example 1: Adding Images")
	pdf.Br(30)

	// Add PNG image
	// Parameters: path, x, y, optional rect (width, height)
	err := pdf.Image("images/photo.jpg", 50, 100, &gopdf.Rect{W: 150, H: 100})
	if err != nil {
		log.Println("Note: Make sure to place logo.png in 'images/' folder")
	}

	// Add JPEG image at different position
	pdf.SetY(220)
	pdf.SetFont("arial", "", 12)
	pdf.Cell(nil, "JPEG image example:")
	err = pdf.Image("images/photo.jpg", 50, 250, &gopdf.Rect{W: 200, H: 150})
	if err != nil {
		log.Println("Note: Make sure to place photo.jpg in /images/ folder")
	}

	// Tips for working with images:
	pdf.SetXY(50, 420)
	pdf.SetFont("arial", "", 10)
	pdf.MultiCell(&gopdf.Rect{W: 500, H: 100},
		"Tips:\n"+
			"- Supported formats: PNG, JPEG\n"+
			"- Use &gopdf.Rect{W: width, H: height} to resize images\n"+
			"- Position with x, y coordinates\n"+
			"- Make sure image files exist in the specified path")

	// Save PDF
	pdf.WritePdf(goPdfFolder + advancedFeatures + "add_images.pdf")

	fmt.Println("Created: add_images.pdf to", goPdfFolder+advancedFeatures, "folder")
}

// Example 2: Drawing Shapes and Lines
func drawShapesExample(pdf *gopdf.GoPdf) {
	pdf.AddPage()

	pdf.SetFont("arial", "", 16)
	pdf.Cell(nil, "Example 2: Drawing Shapes and Lines")
	pdf.Br(30)

	// Set line and fill colors
	pdf.SetLineWidth(2)
	pdf.SetStrokeColor(0, 0, 255) // Blue stroke

	// Draw a rectangle (outline only)
	pdf.SetX(50)
	pdf.SetY(100)
	pdf.RectFromUpperLeftWithStyle(50, 100, 150, 80, "D") // D = Draw (outline only)

	// Draw a filled rectangle
	pdf.SetFillColor(255, 200, 200)                        // Light red fill
	pdf.RectFromUpperLeftWithStyle(250, 100, 150, 80, "F") // F = Fill

	// Draw a rectangle with both fill and outline
	pdf.SetStrokeColor(0, 128, 0)                           // Green stroke
	pdf.SetFillColor(200, 255, 200)                         // Light green fill
	pdf.RectFromUpperLeftWithStyle(450, 100, 100, 80, "FD") // FD = Fill and Draw

	// Draw lines
	pdf.SetLineWidth(1)
	pdf.SetStrokeColor(0, 0, 0) // Black

	// Horizontal line
	pdf.Line(50, 220, 550, 220)

	// Vertical line
	pdf.Line(300, 240, 300, 340)

	// Diagonal line
	pdf.Line(50, 240, 200, 340)

	// Draw a polyline (connected lines)
	pdf.SetLineWidth(2)
	pdf.SetStrokeColor(255, 0, 0) // Red
	// Starting point
	pdf.SetX(400)
	pdf.SetY(240)
	// Draw connected lines
	pdf.Line(400, 240, 450, 280)
	pdf.Line(450, 280, 500, 260)
	pdf.Line(500, 260, 550, 300)

	// Draw a circle (using oval with equal width and height)
	pdf.SetLineWidth(2)
	pdf.SetStrokeColor(0, 0, 255)
	pdf.SetFillColor(200, 200, 255)
	pdf.Oval(150, 380, 50, 50) // x, y, width, height

	// Add labels
	pdf.SetFont("arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(80, 185)
	pdf.Cell(nil, "Outline")
	pdf.SetXY(280, 185)
	pdf.Cell(nil, "Filled")
	pdf.SetXY(480, 185)
	pdf.Cell(nil, "Both")

	// Save PDF
	pdf.WritePdf(goPdfFolder + advancedFeatures + "draw_shapes.pdf")

	fmt.Println("Created: draw_shapes.pdf to", goPdfFolder+advancedFeatures, "folder")
}

// Example 3: Creating Tables and Grids
func createTableExample(pdf *gopdf.GoPdf) {
	pdf.AddPage()

	pdf.SetFont("arial", "", 16)
	pdf.Cell(nil, "Example 3: Creating Tables and Grids")
	pdf.Br(30)

	// Table configuration
	startX := 50.0
	startY := 100.0
	colWidths := []float64{50, 150, 100, 100, 100} // Column widths
	rowHeight := 25.0

	// Table headers
	headers := []string{"ID", "Product Name", "Quantity", "Price", "Total"}

	// Table data
	data := [][]string{
		{"1", "Laptop", "2", "$800", "$1,600"},
		{"2", "Mouse", "5", "$25", "$125"},
		{"3", "Keyboard", "3", "$75", "$225"},
		{"4", "Monitor", "2", "$300", "$600"},
		{"5", "USB Cable", "10", "$5", "$50"},
	}

	// Draw table header
	pdf.SetFillColor(100, 100, 100) // Gray background for header
	pdf.SetTextColor(255, 255, 255) // White text
	pdf.SetFont("arial", "", 11)

	currentY := startY
	currentX := startX

	// Draw header cells
	for i, header := range headers {
		// Draw cell background
		pdf.RectFromUpperLeftWithStyle(currentX, currentY, colWidths[i], rowHeight, "F")

		// Draw cell border
		pdf.SetStrokeColor(0, 0, 0)
		pdf.RectFromUpperLeftWithStyle(currentX, currentY, colWidths[i], rowHeight, "D")

		// Add text
		pdf.SetXY(currentX+5, currentY+8)
		pdf.Cell(nil, header)

		currentX += colWidths[i]
	}

	// Draw table rows
	pdf.SetTextColor(0, 0, 0) // Black text
	pdf.SetFont("arial", "", 10)

	currentY += rowHeight

	for _, row := range data {
		currentX = startX

		for i, cell := range row {
			// Alternate row colors for better readability
			if int((currentY-startY)/rowHeight)%2 == 0 {
				pdf.SetFillColor(240, 240, 240) // Light gray
			} else {
				pdf.SetFillColor(255, 255, 255) // White
			}

			// Draw cell background
			pdf.RectFromUpperLeftWithStyle(currentX, currentY, colWidths[i], rowHeight, "F")

			// Draw cell border
			pdf.SetStrokeColor(0, 0, 0)
			pdf.RectFromUpperLeftWithStyle(currentX, currentY, colWidths[i], rowHeight, "D")

			// Add text
			pdf.SetXY(currentX+5, currentY+8)
			pdf.Cell(nil, cell)

			currentX += colWidths[i]
		}

		currentY += rowHeight
	}

	// Add table summary
	currentY += 20
	pdf.SetFont("arial", "", 12)
	pdf.SetXY(startX, currentY)
	pdf.Cell(nil, "Total Items: 5")
	pdf.SetXY(startX+300, currentY)
	pdf.Cell(nil, "Grand Total: $2,600")

	// Save PDF
	pdf.WritePdf(goPdfFolder + advancedFeatures + "table-create.pdf")

	fmt.Println("Created: table_create.pdf to", goPdfFolder+advancedFeatures, "folder")
}

// Example 4: Headers and Footers Implementation
func addHeaderFooterExample(pdf *gopdf.GoPdf) {
	// Add multiple pages to demonstrate headers and footers
	for i := 1; i <= 3; i++ {
		pdf.AddPage()

		// HEADER
		drawHeader(pdf, "Advanced gopdf Tutorial")

		// CONTENT
		pdf.SetY(80) // Start content below header
		pdf.SetFont("arial", "", 12)
		pdf.Cell(nil, fmt.Sprintf("This is page %d content", i))
		pdf.Br(20)
		pdf.MultiCell(&gopdf.Rect{W: 500, H: 400},
			"Headers and footers are essential for professional documents. "+
				"They provide consistent branding and navigation information across pages.\n\n"+
				"Best practices:\n"+
				"- Keep headers simple and consistent\n"+
				"- Include document title or company name\n"+
				"- Add page numbers in footers\n"+
				"- Include date or document metadata\n"+
				"- Use subtle colors or lines to separate from content")

		// FOOTER
		drawFooter(pdf, i, 3) // Current page, total pages
	}

	// Save PDF
	pdf.WritePdf(goPdfFolder + advancedFeatures + "add_header_footer.pdf")

	fmt.Println("Created: add_header_footer.pdf to", goPdfFolder+advancedFeatures, "folder")
}

// Helper function to draw header
func drawHeader(pdf *gopdf.GoPdf, title string) {
	// Draw header background
	pdf.SetFillColor(41, 128, 185)                     // Blue background
	pdf.RectFromUpperLeftWithStyle(0, 0, 595, 50, "F") // A4 width = 595 points

	// Add header text
	pdf.SetTextColor(255, 255, 255) // White text
	pdf.SetFont("arial", "", 18)
	pdf.SetXY(50, 15)
	pdf.Cell(nil, title)

	// Add header line
	pdf.SetStrokeColor(255, 255, 255)
	pdf.SetLineWidth(2)
	pdf.Line(50, 45, 545, 45)

	// Reset text color for content
	pdf.SetTextColor(0, 0, 0)
}

// Helper function to draw footer
func drawFooter(pdf *gopdf.GoPdf, currentPage, totalPages int) {
	// Draw footer line
	pdf.SetStrokeColor(200, 200, 200)
	pdf.SetLineWidth(1)
	pdf.Line(50, 792, 545, 792) // A4 height = 842 points

	// Add footer text
	pdf.SetTextColor(100, 100, 100) // Gray text
	pdf.SetFont("arial", "", 10)

	// Left side - document info
	pdf.SetXY(50, 800)
	pdf.Cell(nil, "gopdf Tutorial - Advanced Features")

	// Center - date
	pdf.SetXY(250, 800)
	pdf.Cell(nil, "Generated with gopdf")

	// Right side - page number
	pdf.SetXY(500, 800)
	pdf.Cell(nil, fmt.Sprintf("Page %d of %d", currentPage, totalPages))

	// Reset text color
	pdf.SetTextColor(0, 0, 0)
}

// Example 5: Page Numbering with gopdf
func addPageNumberingExample(pdf *gopdf.GoPdf) {
	pdf.AddPage()

	pdf.SetFont("arial", "", 16)
	pdf.Cell(nil, "Example 5: Page Numbering Techniques")
	pdf.Br(30)

	pdf.SetFont("arial", "", 12)
	pdf.MultiCell(&gopdf.Rect{W: 500, H: 400},
		"Page Numbering Techniques:\n\n"+
			"1. Simple Page Numbers:\n"+
			"   - Bottom center: 'Page X'\n"+
			"   - Bottom right: 'X'\n\n"+
			"2. Page X of Y Format:\n"+
			"   - Shows total pages: 'Page 1 of 10'\n"+
			"   - Useful for long documents\n\n"+
			"3. Custom Formats:\n"+
			"   - Chapter-Page: '2-5' (Chapter 2, Page 5)\n"+
			"   - Roman numerals for front matter: 'i, ii, iii'\n"+
			"   - Numbers for main content: '1, 2, 3'\n\n"+
			"4. Implementation Tips:\n"+
			"   - Track page count as you add pages\n"+
			"   - Use helper functions for consistency\n"+
			"   - Consider different numbering for different sections\n"+
			"   - Place numbers consistently (footer or header)\n\n"+
			"Look at the footer of this page for an example!")

	// Add a simple page number in the center bottom
	pdf.SetFont("arial", "", 10)
	pdf.SetXY(275, 800)
	pdf.Cell(nil, "- End of Examples -")

	// Save PDF
	pdf.WritePdf(goPdfFolder + advancedFeatures + "page_numbering.pdf")

	fmt.Println("Created: page_numbering.pdf to", goPdfFolder+advancedFeatures, "folder")
}
