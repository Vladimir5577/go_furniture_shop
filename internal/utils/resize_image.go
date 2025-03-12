package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func ResizeImage(filename string) error {
	path := "./uploads/" + filename
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	imageExtension := filepath.Ext(path)

	// decode jpeg into image.Image
	var img image.Image
	if imageExtension == ".jpeg" || imageExtension == ".jpg" {
		img, err = jpeg.Decode(file)
		if err != nil {
			return err
		}
	}

	if imageExtension == ".png" {
		img, err = png.Decode(file)
		if err != nil {
			return err
		}
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(700, 0, img, resize.Lanczos3)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	switch imageExtension {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(out, m, nil)
		if err != nil {
			return err
		}
	case ".png":
		err = png.Encode(out, m)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported format: %s", imageExtension)
	}

	return nil
}
