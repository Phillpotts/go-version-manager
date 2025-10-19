// Package decompressor provides the ability to create a go version directory
package decompressor

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// func decompressGoArchive(archiveFilePath string, destDir string) (string, error) {
// 	return "", nil
// }

func DecompressSave(compressedFilePath string, destDir string) (string, error) {
	// Open the compressed compressed file
	file, err := os.Open(compressedFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create Gzip reader
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	// Create tar reader
	tarReader := tar.NewReader(gzReader)

	// Extract each file
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("tar read error: %w", err)
		}

		// Construct the full path
		target := filepath.Join(destDir, header.Name)

		// Handle different file types
		switch header.Typeflag {
		case tar.TypeDir:
			// Create directory
			if err := os.Mkdir(target, 0o755); err != nil {
				return "", fmt.Errorf("failed to create a directory: %w", err)
			}

		case tar.TypeReg:
			// Create parent directories if they don't exist
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return "", fmt.Errorf("failed to create parent directory: %w", err)
			}

			// Create the file
			outFile, err := os.Create(target)
			if err != nil {
				return "", fmt.Errorf("failed to create file: %w", err)
			}

			// Copy file content
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return "", fmt.Errorf("failed to write file: %w", err)
			}
			outFile.Close()

		default:
			fmt.Printf("Unsupported typr: %v in %s\n", header.Typeflag, header.Name)
		}
	}

	return "", nil
}
