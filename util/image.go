package util

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"

	"golang.org/x/image/tiff"
)

// signatureInfo holds the byte signature and any special validation flags.
type signatureInfo struct {
	Signature []byte
	IsWebP    bool // True if special WebP validation (checking "WEBP" at offset 8) is required.
}

// imageSignaturesMap maps file extensions to a slice of possible signatures.
// This allows for multiple valid signatures for a single extension (e.g., TIFF).
var imageSignaturesMap = map[string][]signatureInfo{
	".jpg":  {{Signature: []byte{0xFF, 0xD8, 0xFF}}},
	".jpeg": {{Signature: []byte{0xFF, 0xD8, 0xFF}}},
	".png":  {{Signature: []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}}},
	".gif":  {{Signature: []byte{0x47, 0x49, 0x46, 0x38}}},
	".webp": {{Signature: []byte{0x52, 0x49, 0x46, 0x46}, IsWebP: true}}, // RIFF header, followed by "WEBP" at offset 8
	".bmp":  {{Signature: []byte{0x42, 0x4D}}},
	".ico":  {{Signature: []byte{0x00, 0x00, 0x01, 0x00}}},
	".tiff": {
		{Signature: []byte{0x49, 0x49, 0x2A, 0x00}}, // TIFF Little Endian
		{Signature: []byte{0x4D, 0x4D, 0x00, 0x2A}}, // TIFF Big Endian
	},
	".tif": {
		{Signature: []byte{0x49, 0x49, 0x2A, 0x00}}, // TIFF Little Endian
		{Signature: []byte{0x4D, 0x4D, 0x00, 0x2A}}, // TIFF Big Endian
	},
}

// IsValidExtension verifies if the initial bytes match the provided extension.
// If the extension does not require byte validation (like .svg), it returns true.
func IsValidExtension(data []byte, ext string) bool {
	ext = strings.ToLower(ext)

	// SVG is a text-based (XML) format and doesn't have a fixed binary signature.
	// For simplicity, it's considered valid if the extension is ".svg".
	// A more robust check might involve parsing the content for "<svg" or "<?xml".
	if ext == ".svg" {
		return true
	}

	signatures, ok := imageSignaturesMap[ext]
	if !ok {
		return true
	}

	for _, sigInfo := range signatures {
		// Check if the data is long enough for the signature and if it matches the prefix.
		if len(data) >= len(sigInfo.Signature) && bytes.HasPrefix(data, sigInfo.Signature) {

			if !sigInfo.IsWebP {
				return true
			}

			if len(data) < 12 {
				continue
			}

			if bytes.Equal(data[8:12], []byte("WEBP")) {
				return true
			}
		}
	}

	// No matching signature found after checking all possibilities for the extension.
	return false
}

const (
	jpegDefaultQuality int = 90 // quality 90 for maintaining weight/visual balance
)

// StripMetadata receives the bytes of an image and its extension, and receives the bytes clean of metadata (Exif, XMP, etc.)
func StripMetadata(data []byte, ext string) ([]byte, error) {
	ext = strings.ToLower(ext)

	switch ext {
	case ".jpg", ".jpeg":
		return stripJpeg(data)
	case ".png":
		return stripPng(data)
	case ".gif":
		return stripGif(data)
	case ".tiff", ".tif":
		return stripTiff(data)
	case ".webp":
		return data, nil
	case ".svg":
		return data, nil
	case ".bmp", ".ico":
		return data, nil
	default:
		return data, nil
	}
}

// stripJpeg decodes and re-encodes the image.
// jpeg.Encode by default does not include additional metadata from the original.
func stripJpeg(data []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: jpegDefaultQuality})
	return buf.Bytes(), err
}

// stripPng uses the standard encoder that only writes critical chunks.
func stripPng(data []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	return buf.Bytes(), err
}

// stripGif re-codifies the GIF removing application blocks (where metadata lives).
func stripGif(data []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = gif.Encode(buf, img, nil)
	return buf.Bytes(), err
}

// stripTiff is a generic function for formats that can contain complex ICC or Exif profiles.
func stripTiff(data []byte) ([]byte, error) {
	img, err := tiff.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tiff.Encode(buf, img, nil)
	return buf.Bytes(), err
}
