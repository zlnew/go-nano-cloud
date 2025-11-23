package handlers

import (
	"fmt"
	"net/http"
	"strings"
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
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(file)))

	w.WriteHeader(200)
	w.Write(file)
}
