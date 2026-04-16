package rail

import (
	"fmt"

	"github.com/example/logistics-freight-calculation-system/calculator"
	"github.com/example/logistics-freight-calculation-system/pricing"
)

type Strategy struct {
	config pricing.Config
}

func New(config pricing.Config) (*Strategy, error) {
	if config.BaseRate <= 0 {
		return nil, fmt.Errorf("rail base rate must be greater than zero")
	}
	return &Strategy{config: config}, nil
}

func (s *Strategy) TransportType() calculator.TransportType {
	return calculator.RailTransport
}

func (s *Strategy) Calculate(input calculator.FreightInput) (float64, error) {
	size := pricing.DimensionalSize(input.Width, input.Height, input.Length)
	price := s.config.BaseRate +
		(input.Weight * s.config.WeightFactor * 0.85) +
		(input.Volume * s.config.VolumeFactor) +
		(size * s.config.SizeFactor)

	if input.Weight > 5000 {
		price += s.config.OversizePenalty
	}
	return pricing.ApplyMinimum(price, s.config.MinimumPrice), nil
}
