package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"net/http"
)

type CompanyHandler struct {
	usecase *usecase.CompanyUsecase
}

func NewCompanyHandler(uc *usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{usecase: uc}
}

func (h *CompanyHandler) ListCompanies(c *gin.Context) {
	companies, err := h.usecase.ListCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}
