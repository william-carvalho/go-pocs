package builder

import (
	"github.com/example/logistics-freight-calculation-system/calculator"
	"github.com/example/logistics-freight-calculation-system/config"
	"github.com/example/logistics-freight-calculation-system/pricing"
	boatpricing "github.com/example/logistics-freight-calculation-system/pricing/boat"
	railpricing "github.com/example/logistics-freight-calculation-system/pricing/rail"
	truckpricing "github.com/example/logistics-freight-calculation-system/pricing/truck"
)

type FreightCalculatorBuilder struct {
	strategies []calculator.Strategy
	errs       []error
}

func NewFreightCalculatorBuilder() *FreightCalculatorBuilder {
	return &FreightCalculatorBuilder{}
}

func (b *FreightCalculatorBuilder) AddStrategy(strategy calculator.Strategy) *FreightCalculatorBuilder {
	if strategy != nil {
		b.strategies = append(b.strategies, strategy)
	}
	return b
}

func (b *FreightCalculatorBuilder) AddTruckPricing() *FreightCalculatorBuilder {
	return b.AddTruckPricingConfig(config.TruckPricing())
}

func (b *FreightCalculatorBuilder) AddTruckPricingConfig(cfg pricing.Config) *FreightCalculatorBuilder {
	strategy, err := truckpricing.New(cfg)
	if err != nil {
		b.errs = append(b.errs, err)
		return b
	}
	return b.AddStrategy(strategy)
}

func (b *FreightCalculatorBuilder) AddBoatPricing() *FreightCalculatorBuilder {
	return b.AddBoatPricingConfig(config.BoatPricing())
}

func (b *FreightCalculatorBuilder) AddBoatPricingConfig(cfg pricing.Config) *FreightCalculatorBuilder {
	strategy, err := boatpricing.New(cfg)
	if err != nil {
		b.errs = append(b.errs, err)
		return b
	}
	return b.AddStrategy(strategy)
}

func (b *FreightCalculatorBuilder) AddRailPricing() *FreightCalculatorBuilder {
	return b.AddRailPricingConfig(config.RailPricing())
}

func (b *FreightCalculatorBuilder) AddRailPricingConfig(cfg pricing.Config) *FreightCalculatorBuilder {
	strategy, err := railpricing.New(cfg)
	if err != nil {
		b.errs = append(b.errs, err)
		return b
	}
	return b.AddStrategy(strategy)
}

func (b *FreightCalculatorBuilder) Build() (*calculator.Calculator, error) {
	if len(b.errs) > 0 {
		return nil, b.errs[0]
	}
	return calculator.New(b.strategies...)
}
