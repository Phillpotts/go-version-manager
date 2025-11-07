// Package downloader provides the capabilities to download go version
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func GetGoVersion(version string, saveDir string, destFileNamePrefix string) (string, error) {
	baseURL := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
	fileName := fmt.Sprintf("%s%s.tar.gz", destFileNamePrefix, version)
	filepath := path.Join(saveDir, fileName)

	// Ensure base directory exists
	err := os.MkdirAll(saveDir, 0o755)
	if err != nil {
		return "", fmt.Errorf("failed to create base directory: %w", err)
	}

	// Create File to be saved
	out, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Download File
	res, err := http.Get(baseURL)
	if err != nil {
		out.Close()
		cleanVersion(filepath)
		return "", fmt.Errorf("failed to download go version")
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		out.Close()
		cleanVersion(filepath)
		return "", fmt.Errorf("go version %s not found", version)
	}

	// Check Server response
	if res.StatusCode != http.StatusOK {
		out.Close()
		cleanVersion(filepath)
		return "", fmt.Errorf("failed to download go version with status code %d", res.StatusCode)
	}

	// Write Body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		cleanVersion(filepath)
		return "", fmt.Errorf("failed to save go version")
	}

	return filepath, nil
}

func cleanVersion(filepath string) {
	err := os.Remove(filepath)
	fmt.Printf("failed to remove file: %s: %s\n", filepath, err)
}
