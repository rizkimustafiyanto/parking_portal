package service

import "mime/multipart"

type Service interface {
	Save(fileHeader *multipart.FileHeader, uploadType string) (*SavedFile, error)
}

type SavedFile struct {
	FileName string
	FilePath string
	FileURL  string
	Type     string
	MimeType string
	Size     int64
}

