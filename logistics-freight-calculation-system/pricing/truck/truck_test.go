package truck

import (
	"testing"

	"github.com/example/logistics-freight-calculation-system/calculator"
	"github.com/example/logistics-freight-calculation-system/pricing"
)

func TestCalculateTruckPrice(t *testing.T) {
	t.Parallel()

	strategy, err := New(pricing.Config{
		BaseRate:        100,
		WeightFactor:    0.5,
		VolumeFactor:    10,
		SizeFactor:      2,
		MinimumPrice:    120,
		OversizePenalty: 50,
	})
	if err != nil {
		t.Fatalf("new strategy: %v", err)
	}

	price, err := strategy.Calculate(calculator.FreightInput{
		Weight: 100,
		Volume: 5,
		Width:  2,
		Height: 2,
		Length: 2,
	})
	if err != nil {
		t.Fatalf("calculate: %v", err)
	}

	if price <= 0 {
		t.Fatalf("expected positive price, got %f", price)
	}
}
