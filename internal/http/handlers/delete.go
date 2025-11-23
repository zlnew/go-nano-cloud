package handlers

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"go/mini-s3/internal/storage"
)

func (h *StorageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := strings.TrimPrefix(r.URL.Path, "/delete/")
	if filename == "" {
		http.Error(w, "filename required", http.StatusBadRequest)
		return
	}

	err := h.Storage.Delete(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			http.NotFound(w, r)
			return
		}

		if errors.Is(err, storage.ErrInvalidFilename) {
			http.Error(w, "invalid filename", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "deleted:", filename)
}
