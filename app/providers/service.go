package providers

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func ReadFolder(path string) ([]FileInfo, error) {
	var files []FileInfo
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file := FileInfo{
				Name:    info.Name(),
				Size:    info.Size(),
				ModTime: info.ModTime(),
			}
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
