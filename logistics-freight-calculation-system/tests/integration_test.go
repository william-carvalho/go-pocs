package tests

import (
	"testing"

	"github.com/example/logistics-freight-calculation-system/builder"
	"github.com/example/logistics-freight-calculation-system/calculator"
)

func TestBuilderCreatesWorkingCalculator(t *testing.T) {
	t.Parallel()

	freightCalculator, err := builder.NewFreightCalculatorBuilder().
		AddTruckPricing().
		AddBoatPricing().
		AddRailPricing().
		Build()
	if err != nil {
		t.Fatalf("build: %v", err)
	}

	price, err := freightCalculator.Calculate(calculator.FreightInput{
		Weight:        1200,
		Volume:        15.5,
		Width:         2.0,
		Height:        1.8,
		Length:        4.5,
		TransportType: calculator.TruckTransport,
	})
	if err != nil {
		t.Fatalf("calculate: %v", err)
	}
	if price <= 0 {
		t.Fatalf("expected positive price, got %f", price)
	}
}
