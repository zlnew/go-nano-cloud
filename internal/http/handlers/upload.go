package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"go/nano-cloud/internal/storage"
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
		http.Error(w, "filepath required", http.StatusBadRequest)
		return
	}

	dest := strings.TrimSpace(r.PostFormValue("destination"))
	if dest != "" {
		dest = filepath.Clean(dest)

		if dest == "." || dest == string(filepath.Separator) {
			dest = ""
		} else if filepath.IsAbs(dest) || strings.HasPrefix(dest, "..") {
			http.Error(w, "invalid destination", http.StatusBadRequest)
			return
		}
	}

	savePath := header.Filename
	if dest != "" {
		savePath = filepath.Join(dest, header.Filename)
	}

	if err := h.Storage.Save(file, savePath); err != nil {
		if errors.Is(err, storage.ErrInvalidPath) {
			http.Error(w, "invalid destination", http.StatusBadRequest)
			return
		}

		http.Error(w, "upload failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "uploaded:", savePath)
}
