package service

import (
	"errors"
	"strings"

	"tax-system/dto"
	"tax-system/model"
	"tax-system/repository"
)

var ErrInvalidProduct = errors.New("product is required")
var ErrInvalidState = errors.New("state is required")
var ErrInvalidYear = errors.New("year must be greater than zero")
var ErrInvalidTaxPercent = errors.New("taxPercent must be zero or greater")
var ErrInvalidBaseAmount = errors.New("baseAmount must be zero or greater")

type TaxService struct {
	repo repository.TaxRuleRepository
}

func NewTaxService(repo repository.TaxRuleRepository) *TaxService {
	return &TaxService{repo: repo}
}

func (s *TaxService) CreateRule(input dto.CreateTaxRuleRequest) (model.TaxRule, error) {
	rule := model.TaxRule{
		Product:    normalize(input.Product),
		State:      normalize(input.State),
		Year:       input.Year,
		TaxPercent: input.TaxPercent,
	}

	if err := validateRule(rule); err != nil {
		return model.TaxRule{}, err
	}

	if err := s.repo.Create(rule); err != nil {
		return model.TaxRule{}, err
	}

	return rule, nil
}

func (s *TaxService) ListRules() []model.TaxRule {
	return s.repo.List()
}

func (s *TaxService) GetRule(product, state string, year int) (model.TaxRule, error) {
	if strings.TrimSpace(product) == "" {
		return model.TaxRule{}, ErrInvalidProduct
	}
	if strings.TrimSpace(state) == "" {
		return model.TaxRule{}, ErrInvalidState
	}
	if year <= 0 {
		return model.TaxRule{}, ErrInvalidYear
	}

	return s.repo.Get(normalize(product), normalize(state), year)
}

func (s *TaxService) CalculateTax(input dto.TaxCalculationRequest) (dto.TaxCalculationResponse, error) {
	if strings.TrimSpace(input.Product) == "" {
		return dto.TaxCalculationResponse{}, ErrInvalidProduct
	}
	if strings.TrimSpace(input.State) == "" {
		return dto.TaxCalculationResponse{}, ErrInvalidState
	}
	if input.Year <= 0 {
		return dto.TaxCalculationResponse{}, ErrInvalidYear
	}
	if input.BaseAmount < 0 {
		return dto.TaxCalculationResponse{}, ErrInvalidBaseAmount
	}

	rule, err := s.repo.Get(normalize(input.Product), normalize(input.State), input.Year)
	if err != nil {
		return dto.TaxCalculationResponse{}, err
	}

	taxValue := input.BaseAmount * rule.TaxPercent
	totalAmount := input.BaseAmount + taxValue

	return dto.TaxCalculationResponse{
		Product:     rule.Product,
		State:       rule.State,
		Year:        rule.Year,
		BaseAmount:  input.BaseAmount,
		TaxPercent:  rule.TaxPercent,
		TaxValue:    taxValue,
		TotalAmount: totalAmount,
	}, nil
}

func validateRule(rule model.TaxRule) error {
	if rule.Product == "" {
		return ErrInvalidProduct
	}
	if rule.State == "" {
		return ErrInvalidState
	}
	if rule.Year <= 0 {
		return ErrInvalidYear
	}
	if rule.TaxPercent < 0 {
		return ErrInvalidTaxPercent
	}
	return nil
}

func normalize(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}
