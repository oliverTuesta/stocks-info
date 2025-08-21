package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverTuesta/stocks-info/backend/internal/api"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/db"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/store"
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

	companyStore := store.NewCompanyStoreDB(conn)
	companyUsecase := usecase.NewCompanyUsecase(companyStore)
	companyHandler := api.NewCompanyHandler(companyUsecase)

	router := gin.Default()
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/companies", companyHandler.ListCompanies)
	}
	router.Run()
}
