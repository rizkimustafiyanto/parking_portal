package service

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	maxUploadSize = 10 << 20
	maxImageSide  = 1600
	imageQuality  = 80
)

var allowedUploadTypes = map[string]struct{}{
	"profile_user":    {},
	"violation_photo": {},
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Save(fileHeader *multipart.FileHeader, uploadType string) (*SavedFile, error) {
	if _, ok := allowedUploadTypes[uploadType]; !ok {
		return nil, errors.New("invalid upload type")
	}

	if fileHeader == nil {
		return nil, errors.New("file is required")
	}

	if fileHeader.Size > maxUploadSize {
		return nil, fmt.Errorf("file too large, max size is %d MB", maxUploadSize>>20)
	}

	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	content, err := io.ReadAll(io.LimitReader(src, maxUploadSize+1))
	if err != nil {
		return nil, err
	}
	if int64(len(content)) > maxUploadSize {
		return nil, fmt.Errorf("file too large, max size is %d MB", maxUploadSize>>20)
	}

	mimeType := http.DetectContentType(content)
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext == "" {
		ext = extensionFromMime(mimeType)
	}

	if !isAllowedFile(ext, mimeType) {
		return nil, errors.New("only jpg, jpeg, png, and pdf files are allowed")
	}

	if err := os.MkdirAll(filepath.Join("uploads", uploadType), 0o755); err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s_%d%s", strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename)), time.Now().UnixNano(), ext)
	fileName = sanitizeFileName(fileName)
	fullPath := filepath.Join("uploads", uploadType, fileName)

	if isImage(ext, mimeType) {
		resized, outMime, err := resizeImage(content, ext)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(fullPath, resized, 0o644); err != nil {
			return nil, err
		}

		return &SavedFile{
			FileName: fileName,
			FilePath: fullPath,
			FileURL:  "/" + filepath.ToSlash(fullPath),
			Type:     uploadType,
			MimeType: outMime,
			Size:     int64(len(resized)),
		}, nil
	}

	if err := os.WriteFile(fullPath, content, 0o644); err != nil {
		return nil, err
	}

	return &SavedFile{
		FileName: fileName,
		FilePath: fullPath,
		FileURL:  "/" + filepath.ToSlash(fullPath),
		Type:     uploadType,
		MimeType: mimeType,
		Size:     int64(len(content)),
	}, nil
}

func isAllowedFile(ext, mimeType string) bool {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".png", ".pdf":
		return true
	}

	switch mimeType {
	case "image/jpeg", "image/png", "application/pdf":
		return true
	}

	return false
}

func isImage(ext, mimeType string) bool {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".png":
		return true
	}

	switch mimeType {
	case "image/jpeg", "image/png":
		return true
	}

	return false
}

func extensionFromMime(mimeType string) string {
	switch mimeType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "application/pdf":
		return ".pdf"
	default:
		return ""
	}
}

func resizeImage(content []byte, ext string) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, "", err
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	maxSide := width
	if height > maxSide {
		maxSide = height
	}

	if maxSide > maxImageSide {
		newWidth := width * maxImageSide / maxSide
		newHeight := height * maxImageSide / maxSide
		if newWidth < 1 {
			newWidth = 1
		}
		if newHeight < 1 {
			newHeight = 1
		}

		dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
		for y := 0; y < newHeight; y++ {
			srcY := bounds.Min.Y + y*height/newHeight
			for x := 0; x < newWidth; x++ {
				srcX := bounds.Min.X + x*width/newWidth
				dst.Set(x, y, img.At(srcX, srcY))
			}
		}
		img = dst
	}

	return reencodeImage(img, format, ext)
}

func reencodeImage(img image.Image, format, ext string) ([]byte, string, error) {
	var buf bytes.Buffer

	if format == "png" || strings.ToLower(ext) == ".png" {
		if err := png.Encode(&buf, img); err != nil {
			return nil, "", err
		}
		return buf.Bytes(), "image/png", nil
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	if err := jpeg.Encode(&buf, rgba, &jpeg.Options{Quality: imageQuality}); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), "image/jpeg", nil
}

func sanitizeFileName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	name = strings.ReplaceAll(name, "/", "_")
	return name
}

