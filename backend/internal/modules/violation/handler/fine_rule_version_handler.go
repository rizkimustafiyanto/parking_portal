package handler

import (
	"net/http"

	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/service"
	pagedto "backend/pkg/dto"
	"backend/pkg/pagination"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type FineRuleVersionHandler struct {
	service service.FineRuleVersionService
}

func NewFineRuleVersionHandler(service service.FineRuleVersionService) *FineRuleVersionHandler {
	return &FineRuleVersionHandler{service: service}
}

func (h *FineRuleVersionHandler) Create(c *gin.Context) {
	var req dto.CreateFineRuleVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}
	if err := h.service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.Success("fine rule version created", nil))
}

func (h *FineRuleVersionHandler) GetByID(c *gin.Context) {
	data, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("fine rule version found", data))
}

func (h *FineRuleVersionHandler) GetAll(c *gin.Context) {
	var query pagedto.PaginationDTO
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid pagination query"))
		return
	}
	query.Normalize()

	var filter dto.ListFineRuleVersionRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid filter query"))
		return
	}

	data, total, err := h.service.GetAll(query, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	meta := pagination.BuildMeta(query.Page, query.Limit, total)
	c.JSON(http.StatusOK, response.SuccessWithMeta("fine rule versions found", data, meta))
}

func (h *FineRuleVersionHandler) Update(c *gin.Context) {
	var req dto.UpdateFineRuleVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}
	if err := h.service.Update(c.Param("id"), req); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("fine rule version updated", nil))
}

func (h *FineRuleVersionHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("fine rule version deleted", nil))
}
