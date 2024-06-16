package main

import (
	"log"
	"net/http"

	httpAdapter "github.com/AviralDixit-star/exoplanet-service/internals/adapters/http"

	"github.com/AviralDixit-star/exoplanet-service/internals/adapters/persistence"
	"github.com/AviralDixit-star/exoplanet-service/internals/application"
)

func main() {
	repo := persistence.NewInMemoryExoplanetRepository()
	service := application.NewExoplanetService(repo)
	router := httpAdapter.NewRouter(service)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
