package calculator

import (
	"fmt"
	"strings"
)

type TransportType string

const (
	BoatTransport  TransportType = "boat"
	TruckTransport TransportType = "truck"
	RailTransport  TransportType = "rail"
)

type FreightInput struct {
	Weight        float64
	Volume        float64
	Width         float64
	Height        float64
	Length        float64
	TransportType TransportType
}

func (i FreightInput) Normalize() FreightInput {
	i.TransportType = TransportType(strings.ToLower(strings.TrimSpace(string(i.TransportType))))
	return i
}

func (i FreightInput) Validate() error {
	if i.Weight <= 0 {
		return fmt.Errorf("weight must be greater than zero")
	}
	if i.Volume <= 0 {
		return fmt.Errorf("volume must be greater than zero")
	}
	if i.Width <= 0 || i.Height <= 0 || i.Length <= 0 {
		return fmt.Errorf("dimensions must be greater than zero")
	}
	if strings.TrimSpace(string(i.TransportType)) == "" {
		return fmt.Errorf("transport type is required")
	}
	return nil
}
