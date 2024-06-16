// internal/domain/entities.go
package domain

import "errors"

type Exoplanet struct {
	ID          string
	Name        string
	Description string
	Distance    int
	Radius      float64
	Mass        float64
	Type        ExoplanetType
}

func GetExoplanetTypeByName(name string) (ExoplanetType, error) {
	switch name {
	case "GasGiant":
		return &GasGiant{
			ExoType: name,
		}, nil
	case "Terrestrial":
		return &Terrestrial{
			ExoType: name,
		}, nil
	default:
		return nil, errors.New("unknown exoplanet type")
	}
}
