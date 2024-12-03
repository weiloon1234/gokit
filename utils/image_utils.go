package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// decodeImage decodes an image from a multipart file header and returns the image along with its format.
func decodeImage(fileHeader *multipart.FileHeader) (image.Image, string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %v", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	src, format, err := image.Decode(file)
	if err != nil {
		if format == "" {
			_, err := file.Seek(0, 0)
			if err != nil {
				return nil, "", err
			} // Reset file pointer for WebP fallback
			webpImg, webpErr := webp.Decode(file)
			if webpErr == nil {
				return webpImg, "webp", nil
			}
		}
		return nil, "", fmt.Errorf("failed to decode image: %v", err)
	}

	return src, format, nil
}

// encodeImage encodes an image into a buffer using the detected format.
func encodeImage(img image.Image, format string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	switch format {
	case "jpeg":
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to encode JPEG: %v", err)
		}
	case "png":
		err := png.Encode(buf, img)
		if err != nil {
			return nil, fmt.Errorf("failed to encode PNG: %v", err)
		}
	case "bmp":
		err := bmp.Encode(buf, img)
		if err != nil {
			return nil, fmt.Errorf("failed to encode BMP: %v", err)
		}
	case "tiff":
		err := tiff.Encode(buf, img, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to encode TIFF: %v", err)
		}
	case "webp":
		err := webp.Encode(buf, img, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to encode WebP: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	return buf, nil
}

// Resize resizes the image to the specified dimensions without maintaining aspect ratio.
func Resize(fileHeader *multipart.FileHeader, width, height *int) (*bytes.Buffer, error) {
	src, format, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	if width == nil && height == nil {
		return nil, fmt.Errorf("at least one dimension (width or height) must be specified")
	}

	w := 0
	if width != nil {
		w = *width
	}
	h := 0
	if height != nil {
		h = *height
	}

	resized := imaging.Resize(src, w, h, imaging.Lanczos)
	return encodeImage(resized, format)
}

// ResizeDown resizes the image only if the new dimensions are smaller than the original.
func ResizeDown(fileHeader *multipart.FileHeader, width, height *int) (*bytes.Buffer, error) {
	src, format, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	w := origWidth
	h := origHeight

	if width != nil && *width < origWidth {
		w = *width
	}
	if height != nil && *height < origHeight {
		h = *height
	}

	resized := imaging.Resize(src, w, h, imaging.Lanczos)
	return encodeImage(resized, format)
}

// Scale scales the image to fit within the given dimensions, maintaining aspect ratio.
func Scale(fileHeader *multipart.FileHeader, width, height *int) (*bytes.Buffer, error) {
	src, format, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	if width == nil && height == nil {
		return nil, fmt.Errorf("at least one dimension (width or height) must be specified")
	}

	w := 0
	if width != nil {
		w = *width
	}
	h := 0
	if height != nil {
		h = *height
	}

	scaled := imaging.Fit(src, w, h, imaging.Lanczos)
	return encodeImage(scaled, format)
}

// ScaleDown scales the image while maintaining aspect ratio and does not exceed the original size.
func ScaleDown(fileHeader *multipart.FileHeader, width, height *int) (*bytes.Buffer, error) {
	src, _, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	if width != nil && *width > origWidth {
		width = &origWidth
	}
	if height != nil && *height > origHeight {
		height = &origHeight
	}

	return Scale(fileHeader, width, height)
}

// Cover crops the image to fill the specified dimensions, maintaining aspect ratio.
func Cover(fileHeader *multipart.FileHeader, width, height int, position *imaging.Anchor) (*bytes.Buffer, error) {
	src, format, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	// Default to center position if not specified
	pos := imaging.Center
	if position != nil {
		pos = *position
	}

	cropped := imaging.Fill(src, width, height, pos, imaging.Lanczos)
	return encodeImage(cropped, format)
}

// CoverDown scales down and crops the image to fill the dimensions without exceeding the original size.
func CoverDown(fileHeader *multipart.FileHeader, width, height int, position *imaging.Anchor) (*bytes.Buffer, error) {
	src, format, err := decodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	if width > origWidth || height > origHeight {
		return encodeImage(src, format) // Return original if dimensions exceed original size
	}

	return Cover(fileHeader, width, height, position)
}
