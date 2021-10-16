package service

import (
	"log"
	"os"
)

type fileService struct {
}

type FileService interface {
  ReadFile(path string) (*os.File, error)
}

func NewFileService() FileService {
  return &fileService{}
}

// ReadCsv - will try to open file. Will create file instead if not found.
func (fi *fileService) ReadFile(fp string) (*os.File, error) {
  csv, err := os.OpenFile(fp, os.O_CREATE, 0666)

  if err != nil {
    log.Fatalf("fatal error trying to read/create file: %s. \n %s", fp , err)
    return nil, err
  }

  return csv, nil
}