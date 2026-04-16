package pricing

func DimensionalSize(width, height, length float64) float64 {
	return width * height * length
}

func ApplyMinimum(price, minimum float64) float64 {
	if price < minimum {
		return minimum
	}
	return price
}
