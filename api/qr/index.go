package handler

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"joseph-sx/kh-barcode-generator/utils"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const qrSize = 250

func Handler(w http.ResponseWriter, r *http.Request) {
	value := utils.SanitizeQRValue(r.URL.Query().Get("v"))
	if value == "" {
		writeEmptyValueImage(w)
		return
	}

	code, err := qr.Encode(value, qr.M, qr.Auto)
	if err != nil {
		http.Error(w, "Value not supported by QR standard", http.StatusBadRequest)
		return
	}

	scaledCode, err := barcode.Scale(code, qrSize, qrSize)
	if err != nil || scaledCode == nil {
		http.Error(w, "Failed to scale QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=31536000, s-maxage=31536000, immutable")
	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "Display error", http.StatusInternalServerError)
	}
}

func writeEmptyValueImage(w http.ResponseWriter) {
	img := image.NewRGBA(image.Rect(0, 0, qrSize, qrSize))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{254, 226, 226, 255}}, image.Point{}, draw.Src)

	drawer := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: color.RGBA{153, 27, 27, 255}},
		Face: basicfont.Face7x13,
	}
	text := "NO VALUE PROVIDED"
	drawer.Dot = fixed.P((qrSize-drawer.MeasureString(text).Round())/2, (qrSize+13)/2-2)
	drawer.DrawString(text)

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	_ = png.Encode(w, img)
}
