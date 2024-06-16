package domain

type Terrestrial struct {
	ExoType string
}

func (t Terrestrial) CalculateGravity(radius float64, mass float64) float64 {
	return mass / (radius * radius)
}
