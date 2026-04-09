package repository

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"tax-system/model"
)

var ErrTaxRuleAlreadyExists = errors.New("tax rule already exists")
var ErrTaxRuleNotFound = errors.New("tax rule not found")

type TaxRuleRepository interface {
	Create(rule model.TaxRule) error
	List() []model.TaxRule
	Get(product, state string, year int) (model.TaxRule, error)
}

type InMemoryTaxRuleRepository struct {
	mu    sync.RWMutex
	rules map[string]model.TaxRule
}

func NewInMemoryTaxRuleRepository() *InMemoryTaxRuleRepository {
	return &InMemoryTaxRuleRepository{
		rules: make(map[string]model.TaxRule),
	}
}

func (r *InMemoryTaxRuleRepository) Create(rule model.TaxRule) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := buildKey(rule.Product, rule.State, rule.Year)
	if _, exists := r.rules[key]; exists {
		return ErrTaxRuleAlreadyExists
	}

	r.rules[key] = rule
	return nil
}

func (r *InMemoryTaxRuleRepository) List() []model.TaxRule {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rules := make([]model.TaxRule, 0, len(r.rules))
	for _, rule := range r.rules {
		rules = append(rules, rule)
	}
	return rules
}

func (r *InMemoryTaxRuleRepository) Get(product, state string, year int) (model.TaxRule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rule, exists := r.rules[buildKey(product, state, year)]
	if !exists {
		return model.TaxRule{}, ErrTaxRuleNotFound
	}

	return rule, nil
}

func buildKey(product, state string, year int) string {
	return fmt.Sprintf("%s:%s:%d", strings.ToUpper(strings.TrimSpace(product)), strings.ToUpper(strings.TrimSpace(state)), year)
}
