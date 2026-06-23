package handler

import (
	"net/http"
	"strings"

	"backend/internal/modules/invoice/dto"
	"backend/internal/modules/invoice/service"
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
	var req dto.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := h.service.Create(req); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "invalid member_id") ||
			strings.Contains(strings.ToLower(err.Error()), "fine rule") {
			c.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("invoice created", nil))
}

func (h *Handler) GetByID(c *gin.Context) {
	invoice, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("invoice found", invoice))
}

func (h *Handler) GetDetail(c *gin.Context) {
	invoice, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("invoice detail found", invoice))
}

func (h *Handler) GetMemberTransactions(c *gin.Context) {
	invoices, err := h.service.GetByMemberID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("member transactions found", invoices))
}

func (h *Handler) GetMemberBalanceHistory(c *gin.Context) {
	invoices, err := h.service.GetByMemberID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("member balance history found", invoices))
}

func (h *Handler) GetAll(c *gin.Context) {
	var query pagedto.PaginationDTO
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid pagination query"))
		return
	}
	query.Normalize()

	var filter dto.ListInvoiceRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid filter query"))
		return
	}

	invoices, total, err := h.service.GetAll(query, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	meta := pagination.BuildMeta(query.Page, query.Limit, total)
	c.JSON(http.StatusOK, response.SuccessWithMeta("invoices found", invoices, meta))
}

func (h *Handler) Update(c *gin.Context) {
	var req dto.UpdateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := h.service.Update(c.Param("id"), req); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "invalid member_id") {
			c.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("invoice updated", nil))
}

func (h *Handler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("invoice deleted", nil))
}
