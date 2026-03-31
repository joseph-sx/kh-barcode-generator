package handler

import (
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("v")
	if value == "" {
		w.Header().Set("Content-Type", "image/svg+xml")
		svgError := `<svg width="200" height="50" xmlns="http://www.w3.org/2000/svg">
			<rect width="100%" height="100%" fill="#fee2e2"/>
			<text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle" font-family="sans-serif" font-size="14" fill="#991b1b">NO VALUE PROVIDED</text>
		</svg>`

		w.WriteHeader(http.StatusOK) // Algunos navegadores fallan el render si es 400
		w.Write([]byte(svgError))
		return
	}
	bc, err := code128.Encode(value)
	if err != nil {
		http.Error(w, "Error encoding", http.StatusInternalServerError)
		return
	}
	scaledBC, _ := barcode.Scale(bc, 250, 80)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=31536000, s-maxage=31536000, immutable")

	png.Encode(w, scaledBC)
}
