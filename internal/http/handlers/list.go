package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *StorageHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filepaths, err := h.Storage.List()
	if err != nil {
		http.Error(w, "failed to list files", http.StatusInternalServerError)
		return
	}

	resp, err := json.MarshalIndent(filepaths, "", "  ")
	if err != nil {
		http.Error(w, "failed to encode list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
