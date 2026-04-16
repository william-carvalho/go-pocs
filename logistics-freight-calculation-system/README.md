# Logistics Freight Calculation System

`logistics-freight-calculation-system` is a simple and clean Go project that calculates freight prices through one API while keeping transport-specific pricing rules isolated and easy to extend.

## Project Structure

```text
logistics-freight-calculation-system/
├── builder/
├── calculator/
├── config/
├── examples/
├── pricing/
│   ├── boat/
│   ├── rail/
│   └── truck/
├── tests/
├── go.mod
└── README.md
```

## Goals

- one API for freight calculation
- separate pricing strategies per transport type
- dynamic pricing configuration
- easy extension for new transport types and rules
- clean separation between business rules and application flow

## API

```go
calculator, err := builder.NewFreightCalculatorBuilder().
    AddTruckPricing().
    AddBoatPricing().
    AddRailPricing().
    Build()

price, err := calculator.Calculate(calculator.FreightInput{
    Weight:        1200,
    Volume:        15.5,
    Width:         2.0,
    Height:        1.8,
    Length:        4.5,
    TransportType: calculator.TruckTransport,
})
```

## Design

- `calculator` owns the public API and input validation.
- `pricing.Strategy` is the extension point for all transport types.
- `pricing/truck`, `pricing/boat`, and `pricing/rail` implement independent business rules.
- `config` provides default pricing configs that can be replaced at build time.
- `builder` wires the selected strategies into one calculator instance.

## Dynamic Pricing

Prices are configurable through `pricing.Config`, so rates can change without changing the calculator API.

Example:

```go
builder.NewFreightCalculatorBuilder().
    AddTruckPricingConfig(pricing.Config{
        BaseRate:        150,
        WeightFactor:    0.40,
        VolumeFactor:    9.0,
        SizeFactor:      6.0,
        MinimumPrice:    180,
        OversizePenalty: 100,
    }).
    Build()
```

## Extending With A New Transport Type

1. Create a new package under `pricing/yourtransport`.
2. Implement the `pricing.Strategy` interface.
3. Add it through `AddStrategy(...)` or a builder convenience method.

```go
type Strategy interface {
    TransportType() calculator.TransportType
    Calculate(calculator.FreightInput) (float64, error)
}
```

## Running

```bash
go test ./...
go run ./examples
```

## Notes

- each transport strategy uses its own pricing formula
- the current rules are intentionally simple but organized for future growth
- input validation happens before strategy dispatch
- pricing rules stay outside the application-facing calculator flow
