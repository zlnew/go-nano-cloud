// Package storage
package storage

import "io"

type Storage interface {
	Save(file io.Reader, filename string) error
	Open(filename string) ([]byte, error)
	Delete(filename string) error
	List() ([]FileList, error)
}

type FileList struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}
