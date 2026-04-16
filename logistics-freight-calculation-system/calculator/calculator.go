package calculator

import "fmt"

type Calculator struct {
	strategies map[TransportType]Strategy
}

type Strategy interface {
	TransportType() TransportType
	Calculate(FreightInput) (float64, error)
}

func New(strategies ...Strategy) (*Calculator, error) {
	if len(strategies) == 0 {
		return nil, fmt.Errorf("at least one pricing strategy is required")
	}

	registry := make(map[TransportType]Strategy, len(strategies))
	for _, strategy := range strategies {
		if strategy == nil {
			continue
		}
		registry[strategy.TransportType()] = strategy
	}
	if len(registry) == 0 {
		return nil, fmt.Errorf("at least one valid pricing strategy is required")
	}

	return &Calculator{strategies: registry}, nil
}

func (c *Calculator) Calculate(input FreightInput) (float64, error) {
	input = input.Normalize()
	if err := input.Validate(); err != nil {
		return 0, err
	}

	strategy, ok := c.strategies[input.TransportType]
	if !ok {
		return 0, fmt.Errorf("unsupported transport type %q", input.TransportType)
	}
	return strategy.Calculate(input)
}
