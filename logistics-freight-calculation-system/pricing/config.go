package pricing

type Config struct {
	BaseRate        float64
	WeightFactor    float64
	VolumeFactor    float64
	SizeFactor      float64
	MinimumPrice    float64
	DistanceFactor  float64
	OversizePenalty float64
}
