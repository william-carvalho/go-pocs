package config

import "github.com/example/logistics-freight-calculation-system/pricing"

func TruckPricing() pricing.Config {
	return pricing.Config{
		BaseRate:        120,
		WeightFactor:    0.35,
		VolumeFactor:    8.5,
		SizeFactor:      5.5,
		MinimumPrice:    150,
		OversizePenalty: 90,
	}
}

func BoatPricing() pricing.Config {
	return pricing.Config{
		BaseRate:        200,
		WeightFactor:    0.18,
		VolumeFactor:    6.8,
		SizeFactor:      3.2,
		MinimumPrice:    250,
		OversizePenalty: 140,
	}
}

func RailPricing() pricing.Config {
	return pricing.Config{
		BaseRate:        160,
		WeightFactor:    0.22,
		VolumeFactor:    7.2,
		SizeFactor:      4.3,
		MinimumPrice:    210,
		OversizePenalty: 120,
	}
}
