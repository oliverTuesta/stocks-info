package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/api"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/db"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/external"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/repository"
	"github.com/oliverTuesta/stocks-info/backend/internal/middleware"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"log"
)

func main() {
	conn, err := db.ConnectGorm()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	log.Println("DB connection established")

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	ctx := context.Background()
	geminiClient, err := external.NewGeminiClient(ctx, "gemini-2.5-flash")
	if err != nil {
		log.Fatalf("Failed to create AI client: %v", err)
	}

	companyRepository := repository.NewCompanyRepositoryDB(conn)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository, geminiClient)
	companyHandler := api.NewCompanyHandler(companyUsecase)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/companies", companyHandler.ListCompanies)
		apiRoutes.GET("/companies/ticker/:ticker", companyHandler.GetCompanyByTicker)
		apiRoutes.GET("/companies/hot", companyHandler.GetHotCompanies)
		apiRoutes.GET("/companies/ticker/:ticker/ai-analysis", companyHandler.GetCompanyAIAnalysis)
	}

	router.Run()
}
