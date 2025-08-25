package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/api"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/db"
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

	companyRepository := repository.NewCompanyRepositoryDB(conn)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository)
	companyHandler := api.NewCompanyHandler(companyUsecase)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/companies", companyHandler.ListCompanies)
		apiRoutes.GET("/companies/ticker/:ticker", companyHandler.GetCompanyByTicker)
		apiRoutes.GET("/companies/hot", companyHandler.GetHotCompanies)
	}
	router.Run()
}
