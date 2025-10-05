package main

import (
	"fmt"
	"os"
)

func main() {
	// Replace with the path to your file
	files := []string{
		"01-font-management.pdf",
		"02-text-positioning.pdf",
		"03-line-spacing.pdf",
		"04-utf8-characters.pdf",
	}
	//filePath := "go.pdf"

	for _, file := range files {
		err := os.Remove(file) // use os.RemoveAll for directories
		if err != nil {
			fmt.Printf("Error deleting %s: %v\n", file, err)
		} else {
			fmt.Printf("Deleted %s\n", file)
		}
	}
	fmt.Println("File(s) deleted successfully.")
}
