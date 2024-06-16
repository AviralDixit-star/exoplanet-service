package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/AviralDixit-star/exoplanet-service/internals/application"
	"github.com/AviralDixit-star/exoplanet-service/internals/domain"
	"github.com/gorilla/mux"
)

// Validate function to check constraints
func validateExoplanet(exoplanet ExoplanetDTO) error {
	if exoplanet.Distance < 10 || exoplanet.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Mass == 0 || exoplanet.Mass < 0.1 || exoplanet.Mass > 10 {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units")
	}
	if exoplanet.Type != "Terrestrial" && exoplanet.Type != "GasGiant" {
		return errors.New("exo Type should be either Terrestrial or Gas Giant")
	}
	return nil
}

type ExoplanetDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass"`
	Type        string  `json:"type"`
}

func (dto *ExoplanetDTO) ToDomain() (domain.Exoplanet, error) {
	exoplanetType, err := domain.GetExoplanetTypeByName(dto.Type)
	if err != nil {
		return domain.Exoplanet{}, err
	}
	return domain.Exoplanet{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Distance:    dto.Distance,
		Radius:      dto.Radius,
		Mass:        dto.Mass,
		Type:        exoplanetType,
	}, nil
}

func AddExoplanet(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto ExoplanetDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := validateExoplanet(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		exoplanet, err := dto.ToDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := service.AddExoplanet(exoplanet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode("New Exo planet has been created")
		if err != nil {
			log.Fatal(err)
		}

	}
}

func ListExoplanets(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exoplanets, err := service.ListExoplanets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-Type", "application/json")
		json.NewEncoder(w).Encode(exoplanets)
	}
}

func GetExoplanet(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		exoplanet, err := service.GetExoplanetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Add("content-Type", "application/json")
		json.NewEncoder(w).Encode(exoplanet)
	}
}

func UpdateExoplanet(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto ExoplanetDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := validateExoplanet(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		exoplanet, err := dto.ToDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := service.UpdateExoplanet(exoplanet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode("Exo planet has been updated")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteExoplanet(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if err := service.DeleteExoplanet(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("Exo planet has been deleted")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func EstimateFuel(service *application.DefaultExoplanetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		crewCap := r.URL.Query().Get("crew_capacity")
		crewCapacity, err := strconv.Atoi(crewCap)
		if err != nil {
			log.Fatal(err)
			return
		}

		exoplanet, err := service.GetExoplanetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		fuel, err := service.EstimateFuel(exoplanet, crewCapacity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(fuel)
	}
}
