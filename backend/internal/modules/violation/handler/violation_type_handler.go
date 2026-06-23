package handler

import (
	"net/http"
	"strings"

	"backend/internal/middleware"
	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/service"
	pagedto "backend/pkg/dto"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ViolationTypeHandler struct {
	service service.ViolationTypeService
}

func NewViolationTypeHandler(service service.ViolationTypeService) *ViolationTypeHandler {
	return &ViolationTypeHandler{service: service}
}

func (h *ViolationTypeHandler) Create(c *gin.Context) {
	var req dto.CreateViolationTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}
	createdByID, err := uuid.Parse(c.GetString(middleware.ContextUserIDKey))
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error("unauthorized"))
		return
	}
	req.CreatedByID = createdByID
	if err := h.service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.Success("violation type created", nil))
}

func (h *ViolationTypeHandler) GetByID(c *gin.Context) {
	item, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("violation type found", item))
}

func (h *ViolationTypeHandler) GetAll(c *gin.Context) {
	var query pagedto.PaginationDTO
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid pagination query"))
		return
	}
	query.Normalize()

	var filter dto.ListViolationTypeRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid filter query"))
		return
	}

	items, total, err := h.service.GetAll(query, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	meta := pagination.BuildMeta(query.Page, query.Limit, total)
	c.JSON(http.StatusOK, response.SuccessWithMeta("violation types found", items, meta))
}

func (h *ViolationTypeHandler) Update(c *gin.Context) {
	var req dto.UpdateViolationTypeRequest
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
	c.JSON(http.StatusOK, response.Success("violation type updated", nil))
}

func (h *ViolationTypeHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("violation type deleted", nil))
}
