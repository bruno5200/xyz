package util_test

import (
	"testing"

	"github.com/bruno5200/xyz/util"
)

func TestIsValidExtension(t *testing.T) {
	tests := []struct {
		desc       string
		fileHeader []byte
		ext        string
		want       bool
	}{
		{
			desc:       "Valid PNG",
			fileHeader: []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00},
			ext:        ".png",
			want:       true,
		},
		{
			desc:       "Invalid PNG",
			fileHeader: []byte{0x50, 0x89, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00},
			ext:        ".png",
			want:       false,
		},
		{
			desc:       "Valid JPEG",
			fileHeader: []byte{0xFF, 0xD8, 0xFF, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00},
			ext:        ".jpg",
			want:       true,
		},
		{
			desc:       "Invalid JPEG",
			fileHeader: []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, 0x00, 0x00},
			ext:        ".jpg",
			want:       false,
		},
		{
			desc:       "Valid GIF",
			fileHeader: []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00},
			ext:        ".gif",
			want:       true,
		},
		{
			desc:       "Invalid GIF",
			fileHeader: []byte{0x49, 0x47, 0x46, 0x38, 0x37, 0x61, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			ext:        ".gif",
			want:       false,
		},
		{
			desc:       "Valid WebP",
			fileHeader: []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x50},
			ext:        ".webp",
			want:       true,
		},
		{
			desc:       "Invalid WebP",
			fileHeader: []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x00},
			ext:        ".webp",
			want:       false,
		},
		{
			desc:       "Valid BMP",
			fileHeader: []byte{0x42, 0x4D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			ext:        ".bmp",
			want:       true,
		},
		{
			desc:       "Invalid BMP",
			fileHeader: []byte{0x4D, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			ext:        ".bmp",
			want:       false,
		},
		{
			desc:       "Valid ICO",
			fileHeader: []byte{0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00},
			ext:        ".ico",
			want:       true,
		},
		{
			desc:       "Invalid ICO",
			fileHeader: []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			ext:        ".ico",
			want:       false,
		},
		{
			desc:       "Valid TIFF Little Endian",
			fileHeader: []byte{0x49, 0x49, 0x2A, 0x00, 0x00, 0x00, 0x00, 0x00},
			ext:        ".tiff",
			want:       true,
		},
		{
			desc:       "Valid TIFF Big Endian",
			fileHeader: []byte{0x4D, 0x4D, 0x00, 0x2A, 0x00, 0x00, 0x00, 0x00},
			ext:        ".tif",
			want:       true,
		},
		{
			desc:       "Invalid TIFF",
			fileHeader: []byte{0x49, 0x49, 0x00, 0x2A, 0x00, 0x00, 0x00, 0x00},
			ext:        ".tiff",
			want:       false,
		},
		{
			desc:       "Valid SVG",
			fileHeader: []byte{0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x20, 0x76, 0x65, 0x72, 0x73},
			ext:        ".svg",
			want:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			if got := util.IsValidExtension(test.fileHeader, test.ext); got != test.want {
				t.Errorf("IsValidExtension(%v, %s) = %v, want %v", test.fileHeader, test.ext, got, test.want)
			}
		})
	}
}

func TestStripMetadata(t *testing.T) {	tests := []struct {
		desc string
		ext  string
		data []byte
	}{
		{
			desc: "SVG returns original data",
			ext:  ".svg",
			data: []byte("<svg>...</svg>"),
		},
		{
			desc: "WebP returns original data",
			ext:  ".webp",
			data: []byte("RIFF....WEBP"),
		},
		{
			desc: "BMP returns original data",
			ext:  ".bmp",
			data: []byte("BM...."),
		},
		{
			desc: "Unknown extension returns original data",
			ext:  ".heic",
			data: []byte("Unknown data..."),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := util.StripMetadata(test.data, test.ext)
			if err != nil {
				t.Errorf("StripMetadata() error = %v", err)
			}
			if string(got) != string(test.data) {
				t.Errorf("StripMetadata() got = %v, want %v", got, test.data)
			}
		})
	}
}
