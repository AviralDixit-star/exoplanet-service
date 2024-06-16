package application

import (
	"errors"

	"github.com/AviralDixit-star/exoplanet-service/internals/domain"
)

type ExoplanetService interface {
	AddExoplanet(exo domain.Exoplanet) error
	ListExoplanets() ([]domain.Exoplanet, error)
	GetExoplanetByID(id string) (domain.Exoplanet, error)
	UpdateExoplanet(exo domain.Exoplanet) error
	DeleteExoplanet(id string) error
}

type DefaultExoplanetService struct {
	repository domain.ExoplanetRepository
}

func (s *DefaultExoplanetService) AddExoplanet(exo domain.Exoplanet) error {
	return s.repository.Add(exo)

}

func (s *DefaultExoplanetService) ListExoplanets() ([]domain.Exoplanet, error) {
	return s.repository.List()
}

func (s *DefaultExoplanetService) GetExoplanetByID(id string) (domain.Exoplanet, error) {
	return s.repository.GetByID(id)
}

func (s *DefaultExoplanetService) UpdateExoplanet(exo domain.Exoplanet) error {
	return s.repository.Update(exo)
}

func (s *DefaultExoplanetService) DeleteExoplanet(id string) error {
	return s.repository.Delete(id)
}

func (s *DefaultExoplanetService) EstimateFuel(exo domain.Exoplanet, crewCapacity int) (float64, error) {
	gravity := exo.Type.CalculateGravity(exo.Radius, exo.Mass)
	if gravity == 0 {
		return 0, errors.New("invalid gravity calculation")
	}
	fuel := float64(exo.Distance) / (gravity * gravity) * float64(crewCapacity)
	return fuel, nil
}

func NewExoplanetService(repo domain.ExoplanetRepository) *DefaultExoplanetService {
	return &DefaultExoplanetService{repository: repo}
}
