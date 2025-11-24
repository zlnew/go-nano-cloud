// Package storage
package storage

import "io"

type Storage interface {
	Save(file io.Reader, filepath string) error
	Open(filepath string) ([]byte, error)
	Delete(filepath string) error
	List() ([]FileList, error)
}

type FileList struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}
