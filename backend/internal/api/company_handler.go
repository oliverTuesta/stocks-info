package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"net/http"
)

type CompanyHandler struct {
	usecase *usecase.ListCompanies
}

func NewCompanyHandler(uc *usecase.ListCompanies) *CompanyHandler {
	return &CompanyHandler{usecase: uc}
}

func (h *CompanyHandler) ListCompanies(c *gin.Context) {
	companies, err := h.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}
