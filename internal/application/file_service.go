// Package application provides core business logic.
package application

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"
)

// File service error.
var (
	ErrLargeFile = errors.New("file to large")
)

// FileService represent contract that file service should implement.
type FileService interface {
	// CreateCopy creates copy source file and returns path to created file and error.
	CreateCopy(sourceFile multipart.File, sourceName string) (string, error)

	// Remove remove file by path and return error.
	Remove(filePath string) error
}

// fileService serves file processing business logic.
type fileService struct {
	basePath string
	maxSize  int64
	log      *slog.Logger
}

// NewFileService creates new instance of file service.
func NewFileService(basePath string, maxSize int64, log *slog.Logger) *fileService {
	return &fileService{basePath: basePath, maxSize: maxSize, log: log}
}

// CreateCopy creates copy source file and returns path to created file and error.
func (s *fileService) CreateCopy(sourceFile multipart.File, sourceName string) (string, error) {
	const op = "fileService.CreateCopy"

	tempFile, err := os.CreateTemp(s.basePath, fmt.Sprintf("*_%s", sourceName))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer func() {
		if err := tempFile.Close(); err != nil {
			s.log.Error("failed to close file", slog.String("error", err.Error()))
		}
	}()

	limitReader := io.LimitReader(sourceFile, s.maxSize)
	written, err := io.Copy(tempFile, limitReader)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if written == s.maxSize {
		var b [1]byte
		n, err := sourceFile.Read(b[:])
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("%s: %w", op, err)
		}

		if n > 0 {
			return "", fmt.Errorf("%s: %w", op, ErrLargeFile)
		}

	}

	return fmt.Sprintf("%s%s", s.basePath, tempFile.Name()), nil
}

// Remove remove file by path and return error.
func (s *fileService) Remove(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("fileService.Remove: %w", err)
	}
	return nil
}
