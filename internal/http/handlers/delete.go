package handlers

import (
	"fmt"
	"net/http"
	"strings"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "deleted:", filename)
}
