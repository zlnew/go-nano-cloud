package handlers

import (
	"fmt"
	"net/http"
)

func (h *StorageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, h.Env.MaxRequestBodySize)

	if err := r.ParseMultipartForm(h.Env.MaxMultipartMemory); err != nil {
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", http.StatusBadRequest)
		return
	}

	defer file.Close()

	if header == nil || header.Filename == "" {
		http.Error(w, "filename required", http.StatusBadRequest)
		return
	}

	if err := h.Storage.Save(file, header.Filename); err != nil {
		http.Error(w, "upload failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "uploaded:", header.Filename)
}
