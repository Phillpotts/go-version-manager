// Package commands stores handlers
package commands

import (
	"fmt"

	"github.com/phillpotts/go-version-manager/internal/manager"
)

func DownloadVersion(service manager.Manager, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: download <version>\nExample: download 1.21.0")
	}

	version := args[0]
	fmt.Printf("Downloading Go version %s...\n", version)

	if err := service.DownloadVersion(version); err != nil {
		return fmt.Errorf("failed to download version %s: %w", version, err)
	}

	fmt.Printf("Successfully downloaded Go %s\n", version)
	return nil
}

func ExtractVersion(service manager.Manager, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: extract <version>\nExample: download 1.21.0")
	}

	version := args[0]
	fmt.Printf("Extracting Go version %s...\n", version)

	if err := service.ExtractVersion(args[0]); err != nil {
		return fmt.Errorf("failed to extract version %s: %w", version, err)
	}

	fmt.Printf("Successfully extract Go %s\n", version)
	return nil
}
