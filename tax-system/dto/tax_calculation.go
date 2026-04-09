package dto

type TaxCalculationRequest struct {
	Product    string  `json:"product"`
	State      string  `json:"state"`
	Year       int     `json:"year"`
	BaseAmount float64 `json:"baseAmount"`
}

type TaxCalculationResponse struct {
	Product     string  `json:"product"`
	State       string  `json:"state"`
	Year        int     `json:"year"`
	BaseAmount  float64 `json:"baseAmount"`
	TaxPercent  float64 `json:"taxPercent"`
	TaxValue    float64 `json:"taxValue"`
	TotalAmount float64 `json:"totalAmount"`
}
