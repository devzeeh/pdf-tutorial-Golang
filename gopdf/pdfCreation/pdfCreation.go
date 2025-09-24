package main

import (
	"fmt"
	_ "log"
	_ "strings"

	"github.com/signintech/gopdf"
)

var (
	pdfFolder = "gopdfExample/"
	//pdfPageSetup = "example.pdf"
)

func main() {
	// Choose which example to run
	fmt.Println("Running Basic PDF Creation Examples...")

	// Example 1: Simple Hello World PDF (Already completed)
	createHelloWorldPDF()

	// Example 2: Page setup and configuration
	createPageSetupPDF()

	// Example 3: Basic text placement and formatting
	createTextFormattingPDF()

	// Example 4: Multiple pages handling
	createMultiPagePDF()

	fmt.Println("All PDF examples created successfully!")
}

// Example 1: Simple "Hello World" PDF example (Completed)
func createHelloWorldPDF() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Add a page
	pdf.AddPage()

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

	// Set font
	pdf.SetFont(fontName, "", 14)

	// Add text
	pdf.Cell(nil, "Hello World from GoPDF!")

	// Save PDF
	pdf.WritePdf(pdfFolder + "hello-world.pdf")

	fmt.Println("Created: hello-world.pdf to", pdfFolder, "folder")
}

// Example 2: Page setup and configuration
func createPageSetupPDF() {
	fmt.Println("Creating page setup example...")

	pdf := gopdf.GoPdf{}

	// Different page configurations
	config := gopdf.Config{
		PageSize: *gopdf.PageSizeA4, // A4 size: 595.28 x 841.89 points
		Unit:     gopdf.UnitPT,      // Points unit
	}

	pdf.Start(config)

	// Page 1: Portrait A4 with margins
	pdf.AddPage()

	// Set margins (left, top, right, bottom)
	pdf.SetMargins(50, 50, 50, 50)

	// Add font (with proper error handling)
	fontName := "arial"
	err := pdf.AddTTFFont(fontName, "./fonts/arial.ttf")
	if err != nil {
		// Fallback: try to add a built-in or system font
		fmt.Printf("Custom font not found, trying built-in font\n")
		err = pdf.AddTTFFont(fontName, "C:/Windows/Fonts/arial.ttf") // Windows system font path
		if err != nil {
			// Final fallback - use a basic font name that gopdf might have built-in
			fontName = "Times-Roman" // Basic PostScript font
			fmt.Printf("Using fallback font: %s\n", fontName)
		}
	}

	pdf.SetFont(fontName, "", 16)

	// Show page dimensions
	pdf.SetXY(50, 50)
	pdf.Text("Page Configuration Demo")

	pdf.SetXY(50, 80)
	pdf.Text(fmt.Sprintf("Page Size: A4 (%.2f x %.2f points)",
		gopdf.PageSizeA4.W, gopdf.PageSizeA4.H))

	pdf.SetXY(50, 110)
	pdf.Text("Margins: 50pt on all sides")

	// Page 2: Landscape orientation
	pdf.AddPageWithOption(gopdf.PageOption{
		PageSize: &gopdf.Rect{W: gopdf.PageSizeA4.H, H: gopdf.PageSizeA4.W}, // Swap W and H for landscape
	})

	pdf.SetXY(50, 50)
	pdf.Text("Landscape Page Example")

	pdf.SetXY(50, 80)
	pdf.Text("This page is in landscape orientation")

	// Page 3: Custom page size
	customSize := &gopdf.Rect{W: 400, H: 600} // Custom dimensions
	pdf.AddPageWithOption(gopdf.PageOption{PageSize: customSize})

	pdf.SetXY(50, 50)
	pdf.Text("Custom Page Size")

	pdf.SetXY(50, 80)
	pdf.Text(fmt.Sprintf("Size: %.0f x %.0f points", customSize.W, customSize.H))

	pdf.WritePdf(pdfFolder + "page-setup-example.pdf")
	fmt.Println("Created: page-setup-example.pdf to ", pdfFolder, "folder")
}

// Helper function to safely set up font
func setupFont(pdf *gopdf.GoPdf, fontName string) string {
	// Try local fonts directory first
	err := pdf.AddTTFFont(fontName, "./fonts/"+fontName+".ttf")
	if err != nil {
		// Try Windows system fonts
		err = pdf.AddTTFFont(fontName, "C:/Windows/Fonts/"+fontName+".ttf")
		if err != nil {
			// Try alternative Windows system fonts
			err = pdf.AddTTFFont(fontName, "C:/Windows/Fonts/calibri.ttf")
			if err != nil {
				// Final fallback - return a basic font name
				fmt.Printf("All font attempts failed, using Times-Roman\n")
				return "Times-Roman"
			}
		}
	}
	return fontName
}

// Example 3: Basic text placement and formatting
func createTextFormattingPDF() {
	fmt.Println("Creating text formatting example...")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// Setup fonts safely
	regularFont := setupFont(&pdf, "arial")
	boldFont := setupFont(&pdf, "arial-bold")

	// If bold font failed, use regular font
	if boldFont == "Times-Roman" {
		boldFont = regularFont
	}

	// Title
	pdf.SetFont(boldFont, "", 20)
	pdf.SetXY(50, 50)
	pdf.Text("Text Formatting Examples")

	// Different font sizes
	sizes := []float64{8, 10, 12, 14, 16, 18}
	yPos := 100.0

	for _, size := range sizes {
		pdf.SetFont(regularFont, "", size)
		pdf.SetXY(50, yPos)
		pdf.Text(fmt.Sprintf("This text is %.0f points", size))
		yPos += size + 5 // Dynamic spacing based on font size
	}

	// Text positioning examples
	pdf.SetFont(regularFont, "", 12)

	// Left aligned text
	pdf.SetXY(50, yPos+30)
	pdf.Text("Left aligned text (default)")

	// Center aligned text (manual calculation)
	text := "Center aligned text"
	textWidth, _ := pdf.MeasureTextWidth(text)
	pageWidth := gopdf.PageSizeA4.W
	centerX := (pageWidth - textWidth) / 2
	pdf.SetXY(centerX, yPos+50)
	pdf.Text(text)

	// Right aligned text
	rightText := "Right aligned text"
	rightTextWidth, _ := pdf.MeasureTextWidth(rightText)
	rightX := pageWidth - rightTextWidth - 50 // 50pt margin from right
	pdf.SetXY(rightX, yPos+70)
	pdf.Text(rightText)

	// Line spacing example
	pdf.SetFont(regularFont, "", 12)
	pdf.SetXY(50, yPos+120)
	pdf.Text("Line Spacing Examples:")

	lineHeight := 15.0
	pdf.SetXY(50, yPos+140)
	pdf.Text("Line 1 - Normal spacing")
	pdf.SetXY(50, yPos+140+lineHeight)
	pdf.Text("Line 2 - Normal spacing")

	lineHeight = 25.0
	pdf.SetXY(50, yPos+180)
	pdf.Text("Line 1 - Wide spacing")
	pdf.SetXY(50, yPos+180+lineHeight)
	pdf.Text("Line 2 - Wide spacing")

	// Add text wrapping example
	pdf.SetFont(regularFont, "", 12)
	pdf.SetXY(50, yPos+230)
	pdf.Text("Text Wrapping Example:")

	// Use the wrapText function
	longText := "This is a long text that needs to be wrapped across multiple lines to demonstrate text wrapping functionality "
	finally := wrapText(&pdf, longText, 50, yPos+250, 400) // 400pt max width

	// Use finalY to position next content
	pdf.SetXY(50, finally+10)
	pdf.Text("Content continues after wrapped text...")

	pdf.WritePdf(pdfFolder + "text-formatting-example1.pdf")
	fmt.Println("Created: text-formatting-example1.pdf to", pdfFolder, "folder")
}

// Example 4: Multiple pages handling
func createMultiPagePDF() {
	fmt.Println("Creating multiple pages example...")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Setup font safely
	fontName := setupFont(&pdf, "arial")

	// Create 5 pages with different content
	for pageNum := 1; pageNum <= 5; pageNum++ {
		pdf.AddPage()

		// Page header
		pdf.SetFont(fontName, "", 16)
		pdf.SetXY(50, 50)
		pdf.Text(fmt.Sprintf("Page %d of 5", pageNum))

		// Page content
		pdf.SetFont(fontName, "", 12)
		pdf.SetXY(50, 100)
		pdf.Text(fmt.Sprintf("This is the content for page number %d.", pageNum))

		// Add some dynamic content based on page number
		yPos := 130.0
		for i := 1; i <= pageNum*3; i++ {
			pdf.SetXY(50, yPos)
			pdf.Text(fmt.Sprintf("Content line %d on page %d", i, pageNum))
			yPos += 20
		}

		// Page footer
		pdf.SetFont(fontName, "", 10)
		pdf.SetXY(50, 800) // Near bottom of A4 page
		pdf.Text(fmt.Sprintf("Footer - Page %d", pageNum))

		// Page break demonstration
		if pageNum < 5 {
			pdf.SetXY(50, 750)
			pdf.Text("(Page break after this content)")
		}
	}

	// Add a final summary page
	pdf.AddPage()
	pdf.SetFont(fontName, "", 14)
	pdf.SetXY(50, 50)
	pdf.Text("Summary Page")

	pdf.SetFont(fontName, "", 12)
	pdf.SetXY(50, 80)
	pdf.Text("This document demonstrates:")

	summaryItems := []string{
		"• Multiple page creation",
		"• Dynamic content generation",
		"• Page numbering",
		"• Headers and footers",
		"• Content flow across pages",
	}

	yPos := 110.0
	for _, item := range summaryItems {
		pdf.SetXY(50, yPos)
		pdf.Text(item)
		yPos += 25
	}

	pdf.WritePdf(pdfFolder + "multi-page-example.pdf")
	fmt.Println("Created: multi-page-example.pdf to", pdfFolder, "folder")
}

// Helper function to demonstrate advanced text wrapping (bonus)
func wrapText(pdf *gopdf.GoPdf, text string, x, y, maxWidth float64) float64 {
	words := []string{"This", "is", "a", "long", "text", "that", "needs", "to", "be", "wrapped", "across", "multiple", "lines", "to", "demonstrate", "text", "wrapping", "functionality"}

	currentY := y
	currentLine := ""

	for _, word := range words {
		testLine := currentLine + word + " "
		textWidth, _ := pdf.MeasureTextWidth(testLine)

		if textWidth > maxWidth && currentLine != "" {
			// Print current line and start new one
			pdf.SetXY(x, currentY)
			pdf.Text(currentLine)
			currentY += 15 // Line height
			currentLine = word + " "
		} else {
			currentLine = testLine
		}
	}

	// Print last line
	if currentLine != "" {
		pdf.SetXY(x, currentY)
		pdf.Text(currentLine)
		currentY += 15
	}

	return currentY
}
