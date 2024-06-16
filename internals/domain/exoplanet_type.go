package domain

type ExoplanetType interface {
	CalculateGravity(radius float64, mass float64) float64
}
