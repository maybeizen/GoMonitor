package outputs

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"monitor/models"
	"os"
	"path/filepath"
	"strings"
)

type FileOutputHandler struct {
	filePath string
}

func NewFileOutputHandler(filePath string) (*FileOutputHandler, error) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	
	return &FileOutputHandler{
		filePath: filePath,
	}, nil
}

func (f *FileOutputHandler) Write(info *models.SystemInfo) error {
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	isCompressed := strings.HasSuffix(f.filePath, ".gz")
	
	tempPath := f.filePath + ".tmp"
	if err := writeToFile(tempPath, data, isCompressed); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	if err := os.Rename(tempPath, f.filePath); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("failed to rename temporary file: %w", err)
	}

	return nil
}

func (f *FileOutputHandler) Name() string {
	return "File Output: " + f.filePath
}

func writeToFile(path string, data []byte, compress bool) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if compress {
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		
		_, err = gzWriter.Write(data)
		if err != nil {
			return err
		}
		
		return gzWriter.Flush()
	}
	
	_, err = file.Write(data)
	return err
} 