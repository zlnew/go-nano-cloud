package handlers

import (
	"fmt"
	"net/http"
)

func (h *StorageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, _ := r.FormFile("file")
	defer file.Close()

	err := h.Storage.Save(file, header.Filename)
	if err != nil {
		http.Error(w, "upload failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "uploaded:", header.Filename)
}
