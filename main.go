package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
`Usage:
  %s [options] <string>

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	opt_output := flag.String("o", "qrcode.png", "Specify filename to output (should be .PNG).")
	flag.Parse()

	content := flag.Args()[0]

	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create(*opt_output)
	defer file.Close()

	png.Encode(file, qrCode)
}
