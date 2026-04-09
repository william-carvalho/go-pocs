package dto

type CreateTaxRuleRequest struct {
	Product    string  `json:"product"`
	State      string  `json:"state"`
	Year       int     `json:"year"`
	TaxPercent float64 `json:"taxPercent"`
}
