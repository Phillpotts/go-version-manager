// Package downloader provides the capabilities to download go version
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadGoVersion(version string) (string, error) {
	baseURL := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
	goDownloadDir := ".go-bin/downloads"
	fileName := fmt.Sprintf("go%s.linux-amd64.tar.gz", version)

	// Get user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory")
	}
	filepath := path.Join(home, goDownloadDir, fileName)

	// Create File to be saved
	out, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create file")
	}
	defer out.Close()

	// Download File
	res, err := http.Get(baseURL)
	if err != nil {
		return "", fmt.Errorf("filaed to download go version")
	}
	defer res.Body.Close()

	// Check Server response
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download go version with status code %d", res.StatusCode)
	}

	// Write Body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save go version")
	}

	return filepath, nil
}
