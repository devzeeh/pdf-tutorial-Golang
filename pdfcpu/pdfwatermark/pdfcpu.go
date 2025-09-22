package main

import (
    "log"

    "github.com/pdfcpu/pdfcpu/pkg/api"
    "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func main() {
    wm, err := api.TextWatermark(
        "Go Lang #1",      // text
        "scale:0.5, rot:45, op:.3", // ✅ use "scale" instead of "sc"
        true,   // onTop
        false,  // update
        types.POINTS,
    )
    if err != nil {
        log.Fatal(err)
    }

    err = api.AddWatermarksFile(
        "example.pdf",
        "exampleEdited.pdf",
        nil, // all pages
        wm,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
}
