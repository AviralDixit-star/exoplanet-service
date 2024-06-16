package domain

type GasGiant struct {
	ExoType string
}

func (g GasGiant) CalculateGravity(radius float64, mass float64) float64 {
	return 0.5 / (radius * radius)
}
