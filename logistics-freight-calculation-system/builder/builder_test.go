package builder

import (
	"strings"
	"testing"

	"github.com/example/logistics-freight-calculation-system/pricing"
)

func TestBuildRequiresAtLeastOneStrategy(t *testing.T) {
	t.Parallel()

	_, err := NewFreightCalculatorBuilder().Build()
	if err == nil || !strings.Contains(err.Error(), "at least one pricing strategy") {
		t.Fatalf("expected missing strategy error, got %v", err)
	}
}

func TestBuildReturnsStrategyConfigError(t *testing.T) {
	t.Parallel()

	_, err := NewFreightCalculatorBuilder().
		AddTruckPricingConfig(pricing.Config{}).
		Build()
	if err == nil || !strings.Contains(err.Error(), "truck base rate") {
		t.Fatalf("expected truck config error, got %v", err)
	}
}
