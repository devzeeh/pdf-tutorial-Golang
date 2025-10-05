package main

import (
	"fmt"
	"strings"

	"github.com/signintech/gopdf"
)

var (
	goPdfFolder      = "gopdfExample/"     // Base folder for all examples
	pdfCreation      = "pdfCreation/"      // Folder for PDF creation examples
	textHandling     = "textHandling/"     // Folder for text handling examples
	advancedFeatures = "advancedFeatures/" // Folder for advanced features examples
	//pdfPageSetup = "example.pdf"
)

func main() {
	fmt.Println("=== Text Handling with gopdf Tutorial ===")

	// Example 1: Font management and embedding
	fontManagementExample()

	// Example 2: Text positioning and alignment
	textPositioningExample()

	// Example 3: Line spacing and paragraph formatting
	lineSpacingExample()

	// Example 4: UTF-8 and special character support
	utf8Example()

	fmt.Println("\nAll text handling examples completed!")
}

// Example 1: Font Management and Embedding
func fontManagementExample() {
	fmt.Println("Example 1: Font Management and Embedding")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// ============================================
	// FONT MANAGEMENT BASICS
	// ============================================

	// Method 1: Adding TTF fonts from file
	// You need actual font files for this to work
	err := pdf.AddTTFFont("arial", "./fonts/arial.ttf")
	if err != nil {
		// Fallback to system fonts (Windows example)
		err = pdf.AddTTFFont("arial", "C:/Windows/Fonts/arial.ttf")
		if err != nil {
			fmt.Println("Arial font not found, using fallback")
		}
	}

	// Method 2: Adding multiple font styles
	// Regular, Bold, Italic variants
	pdf.AddTTFFont("arial-bold", "C:/Windows/Fonts/arialbd.ttf")
	pdf.AddTTFFont("arial-italic", "C:/Windows/Fonts/ariali.ttf")

	// Method 3: Adding different font families
	pdf.AddTTFFont("times", "C:/Windows/Fonts/times.ttf")
	pdf.AddTTFFont("courier", "C:/Windows/Fonts/cour.ttf")

	// ============================================
	// USING DIFFERENT FONTS
	// ============================================

	// Title
	pdf.SetFont("arial-bold", "", 18)
	pdf.SetXY(50, 50)
	pdf.Text("Font Management Examples")

	yPos := 80.0

	// Example: Arial Regular
	pdf.SetFont("arial", "", 12)
	pdf.SetXY(50, yPos)
	pdf.Text("This is Arial Regular font")
	yPos += 25

	// Example: Arial Bold
	pdf.SetFont("arial-bold", "", 12)
	pdf.SetXY(50, yPos)
	pdf.Text("This is Arial Bold font")
	yPos += 25

	// Example: Arial Italic
	pdf.SetFont("arial-italic", "", 12)
	pdf.SetXY(50, yPos)
	pdf.Text("This is Arial Italic font")
	yPos += 25

	// Example: Times New Roman
	pdf.SetFont("times", "", 12)
	pdf.SetXY(50, yPos)
	pdf.Text("This is Times New Roman font")
	yPos += 25

	// Example: Courier (monospace)
	pdf.SetFont("courier", "", 12)
	pdf.SetXY(50, yPos)
	pdf.Text("This is Courier font (monospace)")
	yPos += 40

	// ============================================
	// FONT SIZES DEMONSTRATION
	// ============================================

	pdf.SetFont("arial-bold", "", 14)
	pdf.SetXY(50, yPos)
	pdf.Text("Different Font Sizes:")
	yPos += 30

	sizes := []float64{8, 10, 12, 14, 16, 18, 20, 24}
	for _, size := range sizes {
		pdf.SetFont("arial", "", size)
		pdf.SetXY(50, yPos)
		pdf.Text(fmt.Sprintf("Font size %.0fpt - The quick brown fox", size))
		yPos += size + 5
	}

	// ============================================
	// FONT EMBEDDING NOTES
	// ============================================

	pdf.AddPage()
	pdf.SetFont("arial", "", 11)

	notes := []string{
		"FONT EMBEDDING TIPS:",
		"",
		"1. Font files (.ttf) must be accessible to your program",
		"2. Use absolute paths or relative paths from your program location",
		"3. Embedded fonts increase PDF file size",
		"4. Always check AddTTFFont() error for missing fonts",
		"5. Have fallback fonts ready for production",
		"",
		"Common font locations:",
		"• Windows: C:/Windows/Fonts/",
		"• macOS: /Library/Fonts/ or /System/Library/Fonts/",
		"• Linux: /usr/share/fonts/ or ~/.fonts/",
	}

	yPos = 50.0
	for _, note := range notes {
		pdf.SetXY(50, yPos)
		if note == "" {
			yPos += 10
			continue
		}

		// Bold for main heading
		if note == "FONT EMBEDDING TIPS:" {
			pdf.SetFont("arial-bold", "", 14)
			pdf.Text(note)
			yPos += 25
		} else {
			pdf.SetFont("arial", "", 11)
			pdf.Text(note)
			yPos += 18
		}
	}

	pdf.WritePdf(goPdfFolder +  textHandling + "01-font-management.pdf")
	fmt.Println("Created: 01-font-management.pdf to",goPdfFolder + textHandling,"folder")
}

// Example 2: Text Positioning and Alignment
func textPositioningExample() {
	fmt.Println("Example 2: Text Positioning and Alignment")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// Setup font
	setupFont(&pdf, "arial")

	// ============================================
	// BASIC POSITIONING WITH SetXY
	// ============================================

	pdf.SetFont("arial", "", 16)
	pdf.SetXY(50, 50)
	pdf.Text("Text Positioning and Alignment")

	// Draw reference grid lines (for visualization)
	drawGrid(&pdf)

	// Absolute positioning
	pdf.SetFont("arial", "", 11)
	pdf.SetXY(100, 120)
	pdf.Text("← Text at X:100, Y:120")

	pdf.SetXY(200, 150)
	pdf.Text("← Text at X:200, Y:150")

	pdf.SetXY(150, 180)
	pdf.Text("← Text at X:150, Y:180")

	// ============================================
	// TEXT ALIGNMENT EXAMPLES
	// ============================================

	yPos := 250.0

	// Left alignment (default)
	pdf.SetXY(50, yPos)
	pdf.Text("LEFT ALIGNED TEXT (default)")
	yPos += 30

	// Center alignment
	text := "CENTER ALIGNED TEXT"
	textWidth, _ := pdf.MeasureTextWidth(text)
	pageWidth := gopdf.PageSizeA4.W
	centerX := (pageWidth - textWidth) / 2
	pdf.SetXY(centerX, yPos)
	pdf.Text(text)
	yPos += 30

	// Right alignment
	rightText := "RIGHT ALIGNED TEXT"
	rightTextWidth, _ := pdf.MeasureTextWidth(rightText)
	rightX := pageWidth - rightTextWidth - 50
	pdf.SetXY(rightX, yPos)
	pdf.Text(rightText)
	yPos += 50

	// ============================================
	// JUSTIFIED TEXT SIMULATION
	// ============================================

	pdf.SetXY(50, yPos)
	pdf.Text("Justified Text Example (stretched to fit width):")
	yPos += 25

	justifiedText := "This text is justified across the page"
	justifyText(&pdf, justifiedText, 50, yPos, 495) // 495 = A4 width - margins

	// ============================================
	// HELPER FUNCTIONS FOR ALIGNMENT
	// ============================================

	pdf.AddPage()
	pdf.SetFont("arial", "", 14)
	pdf.SetXY(50, 50)
	pdf.Text("Using Alignment Helper Functions")

	// Using helper functions
	yPos = 100.0

	alignLeft(&pdf, "Left aligned using helper", 50, yPos)
	yPos += 30

	alignCenter(&pdf, "Center aligned using helper", yPos)
	yPos += 30

	alignRight(&pdf, "Right aligned using helper", 50, yPos)
	yPos += 50

	// Multiple columns example
	pdf.SetFont("arial", "", 11)
	pdf.SetXY(50, yPos)
	pdf.Text("COLUMN 1")

	pdf.SetXY(250, yPos)
	pdf.Text("COLUMN 2")

	pdf.SetXY(450, yPos)
	pdf.Text("COLUMN 3")

	pdf.WritePdf(goPdfFolder +  textHandling + "02-text-positioning.pdf")
	fmt.Println("Created: 02-text-positioning.pdf to",goPdfFolder + textHandling,"folder")
}

// Example 3: Line Spacing and Paragraph Formatting
func lineSpacingExample() {
	fmt.Println("Example 3: Line Spacing and Paragraph Formatting")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	setupFont(&pdf, "arial")

	// ============================================
	// LINE SPACING EXAMPLES
	// ============================================

	pdf.SetFont("arial", "", 16)
	pdf.SetXY(50, 50)
	pdf.Text("Line Spacing and Paragraph Formatting")

	yPos := 90.0

	// Tight line spacing
	pdf.SetFont("arial", "", 11)
	pdf.SetXY(50, yPos)
	pdf.Text("Tight Line Spacing (12pt):")
	yPos += 20

	lineHeight := 12.0
	for i := 1; i <= 5; i++ {
		pdf.SetXY(50, yPos)
		pdf.Text(fmt.Sprintf("Line %d - This is tightly spaced text", i))
		yPos += lineHeight
	}
	yPos += 20

	// Normal line spacing
	pdf.SetXY(50, yPos)
	pdf.Text("Normal Line Spacing (18pt):")
	yPos += 20

	lineHeight = 18.0
	for i := 1; i <= 5; i++ {
		pdf.SetXY(50, yPos)
		pdf.Text(fmt.Sprintf("Line %d - This is normally spaced text", i))
		yPos += lineHeight
	}
	yPos += 20

	// Wide line spacing
	pdf.SetXY(50, yPos)
	pdf.Text("Wide Line Spacing (25pt):")
	yPos += 20

	lineHeight = 25.0
	for i := 1; i <= 5; i++ {
		pdf.SetXY(50, yPos)
		pdf.Text(fmt.Sprintf("Line %d - This is widely spaced text", i))
		yPos += lineHeight
	}

	// ============================================
	// PARAGRAPH FORMATTING
	// ============================================

	pdf.AddPage()
	pdf.SetFont("arial", "", 14)
	pdf.SetXY(50, 50)
	pdf.Text("Paragraph Formatting Examples")

	yPos = 80.0

	// Paragraph 1: No indentation
	pdf.SetFont("arial", "", 11)
	pdf.SetXY(50, yPos)
	pdf.Text("Paragraph 1 (No Indentation):")
	yPos += 20

	paragraph1 := "This is the first paragraph with no indentation. It starts at the left margin and continues normally. This demonstrates basic paragraph formatting without any special indentation."
	yPos = wrapTextWithSpacing(&pdf, paragraph1, 50, yPos, 495, 15)
	yPos += 20

	// Paragraph 2: First-line indentation
	pdf.SetXY(50, yPos)
	pdf.Text("Paragraph 2 (First-line Indentation):")
	yPos += 20

	// Add indentation to first line
	pdf.SetXY(80, yPos) // Indented 30pt
	pdf.Text("This paragraph has first-line indentation. Notice how the first line")
	yPos += 15

	paragraph2Rest := "starts further to the right, creating a traditional paragraph style commonly seen in books and formal documents."
	yPos = wrapTextWithSpacing(&pdf, paragraph2Rest, 50, yPos, 495, 15)
	yPos += 20

	// Paragraph 3: Hanging indentation
	pdf.SetXY(50, yPos)
	pdf.Text("Paragraph 3 (Hanging Indentation):")
	yPos += 20

	pdf.SetXY(50, yPos)
	pdf.Text("1. This is hanging indentation where the first line starts at the margin")
	yPos += 15

	hangingText := "but subsequent lines are indented. This is commonly used in bibliographies and numbered lists."
	yPos = wrapTextWithSpacing(&pdf, hangingText, 70, yPos, 475, 15)
	yPos += 20

	// Paragraph 4: Block quote style
	pdf.SetXY(50, yPos)
	pdf.Text("Paragraph 4 (Block Quote Style):")
	yPos += 20

	blockQuote := "This entire paragraph is indented from both margins, creating a block quote effect. This style is often used for quotations, examples, or to highlight important text within a document."
	yPos = wrapTextWithSpacing(&pdf, blockQuote, 80, yPos, 435, 15) // Indented from left, narrower width
	yPos += 20

	// ============================================
	// SPACING BETWEEN PARAGRAPHS
	// ============================================

	pdf.AddPage()
	pdf.SetFont("arial", "", 14)
	pdf.SetXY(50, 50)
	pdf.Text("Spacing Between Paragraphs")

	yPos = 80.0

	spacings := []float64{10, 20, 30}
	for i, spacing := range spacings {
		pdf.SetFont("arial", "", 11)
		p := fmt.Sprintf("Paragraph %d: This paragraph demonstrates %.0fpt spacing after. Lorem ipsum dolor sit amet, consectetur adipiscing elit.", i+1, spacing)
		yPos = wrapTextWithSpacing(&pdf, p, 50, yPos, 495, 15)
		yPos += spacing // Add the specified spacing after paragraph
	}

	pdf.WritePdf(goPdfFolder +  textHandling + "03-line-spacing.pdf")
	fmt.Println("Created: 03-line-spacing.pdf to",goPdfFolder + textHandling,"folder")
}

// Example 4: UTF-8 and Special Character Support
func utf8Example() {
	fmt.Println("Example 4: UTF-8 and Special Character Support")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// UTF-8 support requires special font handling
	// Use a font that supports Unicode characters
	err := pdf.AddTTFFont("unicode", "C:/Windows/Fonts/arial.ttf")
	if err != nil {
		fmt.Println("Unicode font not found, some characters may not display")
	}

	pdf.SetFont("unicode", "", 16)
	pdf.SetXY(50, 50)
	pdf.Text("UTF-8 and Special Characters")

	yPos := 90.0
	pdf.SetFont("unicode", "", 12)

	// ============================================
	// BASIC SPECIAL CHARACTERS
	// ============================================

	pdf.SetXY(50, yPos)
	pdf.Text("Common Special Characters:")
	yPos += 25

	specialChars := []string{
		`Quotes: "Hello" 'World' guillemets`,
		"Currency: $ (euro) (pound) (yen)",
		"Math: +/- x / approx not-equal less-equal greater-equal",
		"Symbols: (c) (R) (TM) section paragraph",
		"Arrows: left right up down double-arrow",
		"Bullets: dot circle square diamond",
	}

	for _, chars := range specialChars {
		pdf.SetXY(50, yPos)
		pdf.Text(chars)
		yPos += 20
	}

	yPos += 20

	// ============================================
	// MULTILINGUAL TEXT
	// ============================================

	pdf.SetXY(50, yPos)
	pdf.Text("Multilingual Support:")
	yPos += 25

	// Note: These will only work if your font supports these characters
	languages := []string{
		"English: Hello World",
		"Spanish: Hola Mundo (accented characters)",
		"French: Bonjour le monde (accented characters)",
		"German: Hallo Welt (umlaut characters)",
		"Portuguese: Ola Mundo (accented characters)",
	}

	for _, lang := range languages {
		pdf.SetXY(50, yPos)
		pdf.Text(lang)
		yPos += 20
	}

	// ============================================
	// ESCAPED CHARACTERS
	// ============================================

	pdf.AddPage()
	pdf.SetFont("unicode", "", 14)
	pdf.SetXY(50, 50)
	pdf.Text("Working with Escaped Characters")

	yPos = 90.0
	pdf.SetFont("unicode", "", 11)

	pdf.SetXY(50, yPos)
	pdf.Text("Characters that need special handling:")
	yPos += 25

	// Parentheses need escaping in PDF
	pdf.SetXY(50, yPos)
	pdf.Text("Parentheses: (like this) and [brackets]")
	yPos += 20

	// Backslash
	pdf.SetXY(50, yPos)
	pdf.Text("File paths: C:\\Users\\Documents")
	yPos += 20

	// Quotes
	pdf.SetXY(50, yPos)
	pdf.Text(`Single quotes: 'text' and "double quotes"`)
	yPos += 30

	// ============================================
	// UTF-8 ENCODING TIPS
	// ============================================

	pdf.SetXY(50, yPos)
	pdf.Text("UTF-8 Encoding Tips:")
	yPos += 25

	tips := []string{
		"1. Always use fonts that support Unicode (like Arial Unicode MS)",
		"2. Ensure your Go source files are saved as UTF-8",
		"3. Test special characters before deploying",
		"4. Some fonts may not support all Unicode characters",
		"5. Consider fallback fonts for missing glyphs",
	}

	for _, tip := range tips {
		pdf.SetXY(50, yPos)
		pdf.Text(tip)
		yPos += 20
	}

	pdf.WritePdf(goPdfFolder +  textHandling + "04-utf8-characters.pdf")
	fmt.Println("Created: 04-utf8-characters.pdf to", goPdfFolder +  textHandling,"folder")
}

// ============================================
// HELPER FUNCTIONS
// ============================================

// setupFont safely loads a font with fallback
func setupFont(pdf *gopdf.GoPdf, fontName string) {
	err := pdf.AddTTFFont(fontName, "./fonts/"+fontName+".ttf")
	if err != nil {
		err = pdf.AddTTFFont(fontName, "C:/Windows/Fonts/"+fontName+".ttf")
		if err != nil {
			pdf.AddTTFFont(fontName, "C:/Windows/Fonts/arial.ttf")
		}
	}
}

// drawGrid draws a reference grid for visualization
func drawGrid(pdf *gopdf.GoPdf) {
	// Draw light gray grid lines
	pdf.SetLineWidth(0.5)
	pdf.SetStrokeColor(200, 200, 200)

	// Vertical lines every 50pt
	for x := 50.0; x < 550; x += 50 {
		pdf.Line(x, 100, x, 220)
	}

	// Horizontal lines every 30pt
	for y := 100.0; y < 220; y += 30 {
		pdf.Line(50, y, 550, y)
	}
}

// alignLeft aligns text to the left at specified position
func alignLeft(pdf *gopdf.GoPdf, text string, x, y float64) {
	pdf.SetXY(x, y)
	pdf.Text(text)
}

// alignCenter centers text on the page
func alignCenter(pdf *gopdf.GoPdf, text string, y float64) {
	textWidth, _ := pdf.MeasureTextWidth(text)
	pageWidth := gopdf.PageSizeA4.W
	centerX := (pageWidth - textWidth) / 2
	pdf.SetXY(centerX, y)
	pdf.Text(text)
}

// alignRight aligns text to the right
func alignRight(pdf *gopdf.GoPdf, text string, rightMargin, y float64) {
	textWidth, _ := pdf.MeasureTextWidth(text)
	pageWidth := gopdf.PageSizeA4.W
	rightX := pageWidth - textWidth - rightMargin
	pdf.SetXY(rightX, y)
	pdf.Text(text)
}

// justifyText attempts to justify text by adding space between words
func justifyText(pdf *gopdf.GoPdf, text string, x, y, width float64) {
	words := strings.Fields(text)
	if len(words) <= 1 {
		pdf.SetXY(x, y)
		pdf.Text(text)
		return
	}

	// Calculate total width of all words
	totalWordWidth := 0.0
	for _, word := range words {
		w, _ := pdf.MeasureTextWidth(word)
		totalWordWidth += w
	}

	// Calculate space to distribute between words
	totalSpaces := float64(len(words) - 1)
	spaceWidth := (width - totalWordWidth) / totalSpaces

	// Place words with calculated spacing
	currentX := x
	for i, word := range words {
		pdf.SetXY(currentX, y)
		pdf.Text(word)

		wordWidth, _ := pdf.MeasureTextWidth(word)
		currentX += wordWidth

		if i < len(words)-1 {
			currentX += spaceWidth
		}
	}
}

// wrapTextWithSpacing wraps text with custom line spacing
func wrapTextWithSpacing(pdf *gopdf.GoPdf, text string, x, y, maxWidth, lineHeight float64) float64 {
	words := strings.Fields(text)
	currentY := y
	currentLine := ""

	for _, word := range words {
		testLine := currentLine + word + " "
		textWidth, _ := pdf.MeasureTextWidth(testLine)

		if textWidth > maxWidth && currentLine != "" {
			pdf.SetXY(x, currentY)
			pdf.Text(strings.TrimSpace(currentLine))
			currentY += lineHeight
			currentLine = word + " "
		} else {
			currentLine = testLine
		}
	}

	if currentLine != "" {
		pdf.SetXY(x, currentY)
		pdf.Text(strings.TrimSpace(currentLine))
		currentY += lineHeight
	}

	return currentY
}
