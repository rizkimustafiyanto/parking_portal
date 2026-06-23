package dto

type UploadResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileURL  string `json:"file_url"`
	Type     string `json:"type"`
	MimeType string `json:"mime_type"`
	Size     int64  `json:"size"`
}

