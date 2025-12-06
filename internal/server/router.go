package server

import (
	"clean-architecture-go/internal/interface/controller"
	"net/http"
	"strconv"
	"strings"
)

// Fungsi: buat route mapping. Tetap native net/http—tidak pakai framework.
// Register routes and wire controller functions
func RegisterRoutes(c *controller.CategoryController) {
	// /categories -> GET (list) / POST (create)
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			c.GetAll(w, r)
		case http.MethodPost:
			c.Create(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// /categories/{id} -> GET, PUT, DELETE
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		// path example: /categories/12
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			c.GetByID(w, r, id)
		case http.MethodPut:
			c.Update(w, r, id)
		case http.MethodDelete:
			c.Delete(w, r, id)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
