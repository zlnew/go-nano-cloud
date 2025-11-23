// Package storage
package storage

import "io"

type Storage interface {
	Save(file io.Reader, filename string) error
	Open(filename string) ([]byte, error)
	Delete(filename string) error
	List() ([]string, error)
}
