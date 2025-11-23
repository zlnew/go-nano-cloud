package handlers

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"go/mini-s3/internal/storage"
)

func (h *StorageHandler) Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := strings.TrimPrefix(r.URL.Path, "/files/")
	if filename == "" {
		http.Error(w, "filename required", http.StatusBadRequest)
		return
	}

	file, err := h.Storage.Open(filename)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrInvalidFilename):
			http.Error(w, "invalid filename", http.StatusBadRequest)
		case errors.Is(err, fs.ErrNotExist):
			http.NotFound(w, r)
		default:
			http.Error(w, "failed to download file", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file)))

	w.WriteHeader(200)
	w.Write(file)
}
