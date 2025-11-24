package handlers

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go/nano-cloud/internal/storage"
)

func (h *StorageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	filepath := chi.URLParam(r, "*")
	if filepath == "" {
		http.Error(w, "filepath required", http.StatusBadRequest)
		return
	}

	err := h.Storage.Delete(filepath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			http.NotFound(w, r)
			return
		}

		if errors.Is(err, storage.ErrInvalidPath) {
			http.Error(w, "invalid filepath", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "deleted:", filepath)
}
