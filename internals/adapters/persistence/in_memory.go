package persistence

import (
	"errors"
	"sync"

	"github.com/AviralDixit-star/exoplanet-service/internals/domain"
)

type InMemoryExoplanetRepository struct {
	mu         sync.Mutex
	exoplanets map[string]domain.Exoplanet
}

func NewInMemoryExoplanetRepository() *InMemoryExoplanetRepository {
	return &InMemoryExoplanetRepository{
		exoplanets: make(map[string]domain.Exoplanet),
	}
}

func (r *InMemoryExoplanetRepository) Add(exo domain.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.exoplanets[exo.ID] = exo
	return nil
}

func (r *InMemoryExoplanetRepository) List() ([]domain.Exoplanet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []domain.Exoplanet
	for _, exo := range r.exoplanets {
		result = append(result, exo)
	}
	return result, nil
}

func (r *InMemoryExoplanetRepository) GetByID(id string) (domain.Exoplanet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	exo, exists := r.exoplanets[id]
	if !exists {
		return domain.Exoplanet{}, errors.New("exoplanet not found")
	}
	return exo, nil
}

func (r *InMemoryExoplanetRepository) Update(exo domain.Exoplanet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.exoplanets[exo.ID] = exo
	return nil
}

func (r *InMemoryExoplanetRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.exoplanets, id)
	return nil
}
