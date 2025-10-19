package main

import (
	"log"
	"os"
	"path"

	"github.com/phillpotts/go-version-manager/internal/decompressor"
	"github.com/phillpotts/go-version-manager/internal/downloader"
)

const (
	GoVersion         = "1.25.3"
	GoRootDirName     = ".go-bin"
	GoArchiveDirName  = "archive"
	GoVersionsDirName = "versions"
)

func main() {
	// Get user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home directory: %v", err)
	}

	// Download the version of go and store it in the
	archivePath := path.Join(home, GoRootDirName, GoArchiveDirName)
	archiveFilePath, err := downloader.GetGoVersion(GoVersion, archivePath)
	if err != nil {
		log.Fatalf("failed to download go version: %s", GoVersion)
	}

	// Create the go version directory
	goVersionPath := path.Join(home, GoRootDirName, GoVersionsDirName, GoVersion)
	if err = os.MkdirAll(goVersionPath, 0o755); err != nil {
		log.Fatalf("failed to create the version directory: %v", err)
	}

	// Decompress archived version of GO and save the content to the versioned directory
	_, err = decompressor.DecompressSave(archiveFilePath, goVersionPath)
	if err != nil {
		log.Fatalf("failed to unarchive and save: %v", err)
	}
}
