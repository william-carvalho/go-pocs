package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"tax-system/dto"
	"tax-system/repository"
	"tax-system/service"
)

type TaxHandler struct {
	service *service.TaxService
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewTaxHandler(service *service.TaxService) *TaxHandler {
	return &TaxHandler{service: service}
}

func (h *TaxHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tax-rules", h.handleTaxRules)
	mux.HandleFunc("/tax-rules/", h.handleTaxRuleByKey)
	mux.HandleFunc("/tax/calculate", h.handleCalculateTax)
}

func (h *TaxHandler) handleTaxRules(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createTaxRule(w, r)
	case http.MethodGet:
		h.listTaxRules(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Message: "method not allowed"})
	}
}

func (h *TaxHandler) createTaxRule(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateTaxRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid request body"})
		return
	}

	rule, err := h.service.CreateRule(request)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, rule)
}

func (h *TaxHandler) listTaxRules(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.service.ListRules())
}

func (h *TaxHandler) handleTaxRuleByKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Message: "method not allowed"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tax-rules/")
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "path must be /tax-rules/{product}/{state}/{year}"})
		return
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "year must be a valid integer"})
		return
	}

	rule, err := h.service.GetRule(parts[0], parts[1], year)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, rule)
}

func (h *TaxHandler) handleCalculateTax(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Message: "method not allowed"})
		return
	}

	var request dto.TaxCalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: "invalid request body"})
		return
	}

	response, err := h.service.CalculateTax(request)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidProduct),
		errors.Is(err, service.ErrInvalidState),
		errors.Is(err, service.ErrInvalidYear),
		errors.Is(err, service.ErrInvalidTaxPercent),
		errors.Is(err, service.ErrInvalidBaseAmount):
		writeJSON(w, http.StatusBadRequest, errorResponse{Message: err.Error()})
	case errors.Is(err, repository.ErrTaxRuleAlreadyExists):
		writeJSON(w, http.StatusConflict, errorResponse{Message: err.Error()})
	case errors.Is(err, repository.ErrTaxRuleNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Message: err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, errorResponse{Message: "internal server error"})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
