package bitmap

import (
	"errors"
	"image"
)

// Rotate rotates the image by a specified angle (90, 180, 270, -90, -180, -270).
func Rotate(img image.Image, angle int) (image.Image, error) {
	switch angle {
	case 90, -270:
		return rotateRight(img), nil
	case 180, -180:
		return rotate180(img), nil
	case 270, -90:
		return rotateLeft(img), nil
	default:
		return nil, errors.New("unsupported rotation angle, use 90, 180, 270 or their negative equivalents")
	}
}

// rotateRight rotates the image 90 degrees clockwise.
func rotateRight(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			newImg.Set(bounds.Dy()-y-1, x, img.At(x, y))
		}
	}
	return newImg
}

// rotateLeft rotates the image 90 degrees counterclockwise.
func rotateLeft(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			newImg.Set(y, bounds.Dx()-x-1, img.At(x, y))
		}
	}
	return newImg
}

// rotate180 rotates the image 180 degrees.
func rotate180(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			newImg.Set(bounds.Dx()-x-1, bounds.Dy()-y-1, img.At(x, y))
		}
	}
	return newImg
}
