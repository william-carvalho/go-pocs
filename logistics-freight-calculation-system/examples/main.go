package main

import (
	"fmt"
	"log"

	"github.com/example/logistics-freight-calculation-system/builder"
	"github.com/example/logistics-freight-calculation-system/calculator"
)

func main() {
	freightCalculator, err := builder.NewFreightCalculatorBuilder().
		AddTruckPricing().
		AddBoatPricing().
		AddRailPricing().
		Build()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	fmt.Printf("Freight price: %.2f\n", price)
}
