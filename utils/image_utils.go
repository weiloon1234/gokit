package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"

	"github.com/disintegration/imaging"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

type ProcessingMethod string

const (
	MethodScale      ProcessingMethod = "Scale"
	MethodScaleDown  ProcessingMethod = "ScaleDown"
	MethodResize     ProcessingMethod = "Resize"
	MethodResizeDown ProcessingMethod = "ResizeDown"
	MethodCover      ProcessingMethod = "Cover"
	MethodCoverDown  ProcessingMethod = "CoverDown"
)

type ImageProcessor struct {
	Width    *int             // Target width (optional)
	Height   *int             // Target height (optional)
	Position imaging.Anchor   // Default for Cover and CoverDown
	Method   ProcessingMethod // Specifies the processing method (Scale, ScaleDown, etc.)
}

// NewImageProcessor creates a new instance of ImageProcessor with default settings.
func NewImageProcessor(width, height *int, position imaging.Anchor, method ProcessingMethod) *ImageProcessor {
	return &ImageProcessor{
		Width:    width,
		Height:   height,
		Position: position,
		Method:   method,
	}
}

// DecodeImage decodes an image from a multipart file header.
func DecodeImage(fileHeader *multipart.FileHeader) (image.Image, string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	src, format, err := image.Decode(file)
	if err != nil {
		if format == "" {
			file.Seek(0, 0) // Reset file pointer for WebP fallback
			webpImg, webpErr := webp.Decode(file, &decoder.Options{})
			if webpErr == nil {
				return webpImg, "webp", nil
			}
		}
		return nil, "", fmt.Errorf("failed to decode image: %v", err)
	}

	return src, format, nil
}

// EncodeImage encodes an image into a buffer using the detected format.
func EncodeImage(img image.Image, format string) (*bytes.Buffer, error) {
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
		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		if err != nil {
			return nil, fmt.Errorf("failed to create WebP encoder options: %v", err)
		}
		err = webp.Encode(buf, img, options)
		if err != nil {
			return nil, fmt.Errorf("failed to encode WebP: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	return buf, nil
}

// Resize resizes the image without maintaining aspect ratio.
func (ipc *ImageProcessor) Resize(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	w, h := 0, 0
	if ipc.Width != nil {
		w = *ipc.Width
	}
	if ipc.Height != nil {
		h = *ipc.Height
	}

	resized := imaging.Resize(src, w, h, imaging.Lanczos)
	return EncodeImage(resized, format)
}

// ResizeDown resizes the image only if the new dimensions are smaller.
func (ipc *ImageProcessor) ResizeDown(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth, origHeight := bounds.Dx(), bounds.Dy()

	w, h := origWidth, origHeight
	if ipc.Width != nil && *ipc.Width < origWidth {
		w = *ipc.Width
	}
	if ipc.Height != nil && *ipc.Height < origHeight {
		h = *ipc.Height
	}

	resized := imaging.Resize(src, w, h, imaging.Lanczos)
	return EncodeImage(resized, format)
}

// Scale scales the image while maintaining aspect ratio.
func (ipc *ImageProcessor) Scale(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	w, h := 0, 0
	if ipc.Width != nil {
		w = *ipc.Width
	}
	if ipc.Height != nil {
		h = *ipc.Height
	}

	scaled := imaging.Fit(src, w, h, imaging.Lanczos)
	return EncodeImage(scaled, format)
}

// ScaleDown scales the image while maintaining aspect ratio, not exceeding the original size.
func (ipc *ImageProcessor) ScaleDown(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth, origHeight := bounds.Dx(), bounds.Dy()

	w, h := origWidth, origHeight
	if ipc.Width != nil && *ipc.Width < origWidth {
		w = *ipc.Width
	}
	if ipc.Height != nil && *ipc.Height < origHeight {
		h = *ipc.Height
	}

	scaled := imaging.Fit(src, w, h, imaging.Lanczos)
	return EncodeImage(scaled, format)
}

// Cover crops the image to fit the specified dimensions.
func (ipc *ImageProcessor) Cover(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	cropped := imaging.Fill(src, *ipc.Width, *ipc.Height, ipc.Position, imaging.Lanczos)
	return EncodeImage(cropped, format)
}

// CoverDown scales down and crops the image without exceeding original size.
func (ipc *ImageProcessor) CoverDown(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	src, format, err := DecodeImage(fileHeader)
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	origWidth, origHeight := bounds.Dx(), bounds.Dy()

	if *ipc.Width > origWidth || *ipc.Height > origHeight {
		return EncodeImage(src, format) // Return original if dimensions exceed original size
	}

	cropped := imaging.Fill(src, *ipc.Width, *ipc.Height, ipc.Position, imaging.Lanczos)
	return EncodeImage(cropped, format)
}

// Process determines the appropriate method to call based on the `Method` in the config.
func (ipc *ImageProcessor) Process(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	switch ipc.Method {
	case MethodScale:
		return ipc.Scale(fileHeader)
	case MethodScaleDown:
		return ipc.ScaleDown(fileHeader)
	case MethodResize:
		return ipc.Resize(fileHeader)
	case MethodResizeDown:
		return ipc.ResizeDown(fileHeader)
	case MethodCover:
		return ipc.Cover(fileHeader)
	case MethodCoverDown:
		return ipc.CoverDown(fileHeader)
	default:
		return nil, fmt.Errorf("unsupported processing method: %s", ipc.Method)
	}
}
