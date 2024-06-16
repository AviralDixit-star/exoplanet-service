package http

import (
	"github.com/AviralDixit-star/exoplanet-service/internals/application"
	"github.com/gorilla/mux"
)

func NewRouter(service *application.DefaultExoplanetService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", AddExoplanet(service)).Methods("POST")
	r.HandleFunc("/exoplanets", ListExoplanets(service)).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", GetExoplanet(service)).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", UpdateExoplanet(service)).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", DeleteExoplanet(service)).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", EstimateFuel(service)).Methods("GET")
	return r
}
