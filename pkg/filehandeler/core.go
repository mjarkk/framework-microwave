package filehandeler

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// OpenFile opens a file from path
func OpenFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// PathDoesNotExist checks if a input path exist and return a `true` if so
func PathDoesNotExist(path string) bool {
	if string(path[0]) != "/" {
		dir, err := filepath.Abs(".")
		if err != nil {
			return true
		}
		path = filepath.Join(dir, path)
	}
	stats, err := os.Stat(path)
	if err != nil || !stats.IsDir() {
		return true
	}
	return false
}

// ListFiles returns a list of files in a dir
func ListFiles(path string) ([]string, error) {
	files := []string{}
	if PathDoesNotExist(path) {
		return files, errors.New("Path not valid")
	}
	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range filesInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

// CurrentDir returns the currentdir plus a path to someware
func CurrentDir(reslovePath ...string) string {
	dir, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	return dir
	// baseDir := []string{dir}
	// return filepath.Join(append(baseDir, reslovePath...))
}
