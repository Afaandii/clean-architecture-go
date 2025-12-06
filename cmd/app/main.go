package main

import (
	"clean-architecture-go/internal/domain/usecase"
	"clean-architecture-go/internal/interface/controller"
	"clean-architecture-go/internal/server"
	"fmt"

	"clean-architecture-go/internal/infrastructure/db"
	repo "clean-architecture-go/internal/infrastructure/repository"

	"log"
	"net/http"
	"os"
)

// Fungsi: bootstrap aplikasi — connect DB, wire dependencies, start server.
func main() {
	// connect db
	dbConn, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal("db connect:", err)
	}
	defer dbConn.Close()

	// wire repository, usecase, controller
	categoryRepo := repo.NewCategoryPGRepository(dbConn)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := controller.NewCategoryController(categoryUC)

	// register routes
	server.RegisterRoutes(categoryCtrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
