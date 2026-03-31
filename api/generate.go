package handler

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"joseph-sx/kh-barcode-generator/utils"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rawV := r.URL.Query().Get("v")
	cleanV := utils.SanitizeBarcodeValue(rawV)

	if cleanV == "" {
		width := 250
		height := 80
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		backgroundColor := color.RGBA{254, 226, 226, 255}

		textColor := color.RGBA{153, 27, 27, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
		d := &font.Drawer{
			Dst:  img,
			Src:  &image.Uniform{textColor},
			Face: basicfont.Face7x13,
		}
		text := "NO VALUE PROVIDED"
		textWidth := d.MeasureString(text).Round()
		textHeight := 13
		x := (width / 2) - (textWidth / 2)
		y := (height / 2) + (textHeight / 2) - 2
		d.Dot = fixed.P(x, y)
		d.DrawString(text)

		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		png.Encode(w, img)
		return
	}

	bc, err := code128.Encode(cleanV)

	if err != nil {
		http.Error(w, "Value not supported by barcode standard", http.StatusBadRequest)
		return
	}
	scaledBC, err := barcode.Scale(bc, 250, 80)
	if err != nil || scaledBC == nil {
		http.Error(w, "Failed to scale barcode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=31536000, s-maxage=31536000, immutable")
	err = png.Encode(w, scaledBC)
	if err != nil {
		http.Error(w, "Display error", http.StatusInternalServerError)
	}
}
