package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// Input PDF file path
var (
	pdfFolder = "gopdfExample/"
	pdfFile   = "example.pdf"
)

// pdfcpu watermarks example
var pdfcpuFolder = "pdfcpuExample/"
var pdfcpuFile = "watermarked.pdf"

func main() {
	wm, err := api.TextWatermark(
		"Watermark ko 'to",         // text
		"scale:0.5, rot:45, op:.3", // properties
		true,                       // onTop
		false,                      // update
		types.POINTS,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = api.AddWatermarksFile(
		pdfFolder+pdfFile,       // input file
		pdfcpuFolder+pdfcpuFile, // output file
		nil,                     // all pages
		wm,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watermark added successfully to", pdfcpuFolder+"exampleEdited.pdf")
}
