package storage

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basepath string) *LocalStorage {
	return &LocalStorage{BasePath: basepath}
}

func (l *LocalStorage) Save(file io.Reader, filename string) error {
	fullPath := filepath.Join(l.BasePath, filename)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return err
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

func (l *LocalStorage) Open(filename string) ([]byte, error) {
	fullPath := filepath.Join(l.BasePath, filename)
	return os.ReadFile(fullPath)
}

func (l *LocalStorage) Delete(filename string) error {
	fullPath := filepath.Join(l.BasePath, filename)
	return os.Remove(fullPath)
}

func (l *LocalStorage) List() ([]string, error) {
	var files []string

	err := filepath.Walk(l.BasePath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, strings.TrimPrefix(path, fmt.Sprintf("%v/", l.BasePath)))
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return files, nil
}

