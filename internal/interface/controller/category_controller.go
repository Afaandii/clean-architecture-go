package controller

import (
	"clean-architecture-go/internal/domain/usecase"
	"encoding/json"
	"net/http"
	"strings"
)

// Fungsi: translate HTTP ↔ usecase. Controller tidak melakukan query DB langsung.
type CategoryController struct {
	UC usecase.CategoryUsecase
}

func NewCategoryController(uc usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{UC: uc}
}

// helper: write JSON response
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// GET /categories
func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.UC.GetAll()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

// GET /categories/{id}
func (c *CategoryController) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	category, err := c.UC.GetByID(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if category == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, category)
}

// POST /categories
// body: { "name": "T-Shirt" }
func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	created, err := c.UC.Create(strings.TrimSpace(payload.Name), strings.TrimSpace(payload.Description))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// PUT /categories/{id}
func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request, id int) {
	var payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	updated, err := c.UC.Update(id, strings.TrimSpace(payload.Name), strings.TrimSpace(payload.Description))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// DELETE /categories/{id}
func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := c.UC.Delete(id); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}
