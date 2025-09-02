package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"net/http"
	"strconv"
)

type CompanyHandler struct {
	usecase *usecase.CompanyUsecase
}

func NewCompanyHandler(uc *usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{usecase: uc}
}

func (h *CompanyHandler) ListCompanies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	search := c.Query("search")
	if pageSize > 100 {
		pageSize = 100
	}
	req := domain.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}
	response, err := h.usecase.ListCompanies(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *CompanyHandler) GetCompanyByTicker(c *gin.Context) {
	ticker := c.Param("ticker")
	company, err := h.usecase.GetCompanyByTicker(ticker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) GetHotCompanies(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit > 50 {
		limit = 50
	}
	companies, err := h.usecase.GetHotCompanies(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) GetCompanyAIAnalysis(c *gin.Context) {
	ticker := c.Param("ticker")
	analysis, err := h.usecase.GetCompanyAIAnalysis(ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, analysis)
}
