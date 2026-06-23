package handler

import (
	"net/http"
	"strings"

	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/service"
	pagedto "backend/pkg/dto"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateViolationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := h.service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("violation created", nil))
}

func (h *Handler) GetByID(c *gin.Context) {
	violation, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("violation found", violation))
}

func (h *Handler) GetAll(c *gin.Context) {
	var query pagedto.PaginationDTO
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid pagination query"))
		return
	}
	query.Normalize()

	var filter dto.ListViolationRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid filter query"))
		return
	}

	violations, total, err := h.service.GetAll(query, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	meta := pagination.BuildMeta(query.Page, query.Limit, total)
	c.JSON(http.StatusOK, response.SuccessWithMeta("violations found", violations, meta))
}

func (h *Handler) Update(c *gin.Context) {
	var req dto.UpdateViolationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := h.service.Update(c.Param("id"), req); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "immutable") {
			c.JSON(http.StatusConflict, response.Error(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("violation updated", nil))
}

func (h *Handler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("violation deleted", nil))
}
