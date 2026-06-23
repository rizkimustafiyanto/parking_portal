package handler

import (
	"net/http"

	"backend/internal/modules/upload/dto"
	"backend/internal/modules/upload/service"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Upload(c *gin.Context) {
	const maxUploadSize = 10 << 20
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	uploadType := c.PostForm("type")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("file is required"))
		return
	}

	saved, err := h.service.Save(fileHeader, uploadType)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("file uploaded", dto.UploadResponse{
		FileName: saved.FileName,
		FilePath: saved.FilePath,
		FileURL:  saved.FileURL,
		Type:     saved.Type,
		MimeType: saved.MimeType,
		Size:     saved.Size,
	}))
}

