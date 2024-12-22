package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/atotto/clipboard"
)

func die(s string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", s, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", s)
	}
	os.Exit(1)
}

const Period = 30

// generate an ephemeral passcode, print it on the terminal,
// and also paste it to the clipboard if possible
func generatePasscode(secret string) {
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    Period,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		die("Can't generate passcode from TOTP secret", err)
	}

	clipboard.WriteAll(passcode)
	fmt.Println(passcode)
}

func main() {
	persist := false
	flag.BoolVar(&persist, "persist", false, "keep running and generate codes in a loop")
	flag.Parse()

	fn := flag.Arg(0)
	if fn == "" {
		die("Usage: totp1xun [-persist] IMAGENAME", nil)
	}

	// read the image file
	file, err := os.Open(fn)
	if err != nil {
		die("Can't open QR code image", err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		die("Can't decode QR code image", err)
	}

	// decode the QR code from the image
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	qrReader := qrcode.NewQRCodeReader()
	qrResult, err := qrReader.Decode(bmp, nil)
	if err != nil {
		die("Can't parse QR code from image", err)
	}

	// parse the string from the QR code as a TOTP URL
	key, err := otp.NewKeyFromURL(qrResult.String())
	if err != nil {
		die("QR code doesn't contain a TOTP secret", err)
	}

	// always generate one code
	generatePasscode(key.Secret())

	// ...and in normal mode, just exit after one code
	if !persist {
		os.Exit(0)
	}

	// in persist mode we want to generate a new code every 30 seconds.
	// To maximize the validity interval of these codes, we first sleep
	// until the next multiple of 30 seconds and only then start the loop.
	dt := time.Duration(Period - time.Now().Unix() % Period) * time.Second
	time.Sleep(dt)

	c := time.Tick(Period * time.Second)
	for range c {
		generatePasscode(key.Secret())
	}
}
