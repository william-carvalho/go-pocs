package calculator

import (
	"strings"
	"testing"
)

type stubStrategy struct {
	transportType TransportType
	price         float64
}

func (s stubStrategy) TransportType() TransportType {
	return s.transportType
}

func (s stubStrategy) Calculate(FreightInput) (float64, error) {
	return s.price, nil
}

func TestCalculateUsesMatchingStrategy(t *testing.T) {
	t.Parallel()

	calculatorInstance, err := New(
		stubStrategy{transportType: TruckTransport, price: 100},
		stubStrategy{transportType: BoatTransport, price: 200},
	)
	if err != nil {
		t.Fatalf("new calculator: %v", err)
	}

	price, err := calculatorInstance.Calculate(FreightInput{
		Weight:        1,
		Volume:        1,
		Width:         1,
		Height:        1,
		Length:        1,
		TransportType: TruckTransport,
	})
	if err != nil {
		t.Fatalf("calculate: %v", err)
	}
	if price != 100 {
		t.Fatalf("expected 100, got %f", price)
	}
}

func TestCalculateRejectsUnsupportedTransport(t *testing.T) {
	t.Parallel()

	calculatorInstance, err := New(stubStrategy{transportType: TruckTransport, price: 100})
	if err != nil {
		t.Fatalf("new calculator: %v", err)
	}

	_, err = calculatorInstance.Calculate(FreightInput{
		Weight:        1,
		Volume:        1,
		Width:         1,
		Height:        1,
		Length:        1,
		TransportType: "plane",
	})
	if err == nil || !strings.Contains(err.Error(), "unsupported transport type") {
		t.Fatalf("expected unsupported transport error, got %v", err)
	}
}

var _ Strategy = stubStrategy{}
