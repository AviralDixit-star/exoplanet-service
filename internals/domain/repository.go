package domain

type ExoplanetRepository interface {
	Add(exoplanet Exoplanet) error
	List() ([]Exoplanet, error)
	GetByID(id string) (Exoplanet, error)
	Update(exoplanet Exoplanet) error
	Delete(id string) error
}
