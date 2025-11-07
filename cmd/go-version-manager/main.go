package main

import (
	"fmt"
	"os"
	"path"

	"github.com/phillpotts/go-version-manager/internal/argparser"
	"github.com/phillpotts/go-version-manager/internal/decompressor"
	"github.com/phillpotts/go-version-manager/internal/downloader"
)

const (
	GoRootDirName              = ".go-bin"
	GoArchiveDirName           = "archive"
	GoVersionsDirName          = "versions"
	GoVersionArchiveNamePrefix = "go-version-"
)

func main() {
	argparser, err := buildArgparser()
	if err != nil {
		fmt.Printf("failed to build argsparser")
		os.Exit(1)
	}
	err = argparser.Parse()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func downloadVersion(args []string) error {
	// Get user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	// Download the version of go and store it in the
	archivePath := path.Join(home, GoRootDirName, GoArchiveDirName)
	_, err = downloader.GetGoVersion(args[0], archivePath, GoVersionArchiveNamePrefix)
	return err
}

func extractVersion(args []string) error {
	// Get user home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	// Create the go version directory
	goVersionPath := path.Join(home, GoRootDirName, GoVersionsDirName, args[0])
	if err = os.MkdirAll(goVersionPath, 0o755); err != nil {
		return fmt.Errorf("failed to create the version directory: %w", err)
	}
	// Decompress archived version of GO and save the content to the versioned directory
	archiveFileName := fmt.Sprintf("%s%s.tar.gz", GoVersionArchiveNamePrefix, args[0])
	archiveFilePath := path.Join(home, GoRootDirName, GoArchiveDirName, archiveFileName)
	_, err = decompressor.DecompressSave(archiveFilePath, goVersionPath)
	return err
}

func buildArgparser() (argparser.ArgParser, error) {
	// Build argsparser
	app := argparser.NewArgParser("Go Version Manager", "A Cli to manage locally installed versions of the Go Language", "0.1.0-alpha")
	app.AddCommand("download", "download and archive the passed version", downloadVersion)
	app.AddCommand("extract", "extract an archive ready for use", extractVersion)
	return *app, nil
}
