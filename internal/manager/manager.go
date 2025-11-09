// Package manager provides dry functionality around cli commands.
package manager

import (
	"fmt"
	"os"
	"path"

	"github.com/phillpotts/go-version-manager/internal/config"
	"github.com/phillpotts/go-version-manager/internal/decompressor"
	"github.com/phillpotts/go-version-manager/internal/downloader"
)

type Manager struct {
	homeDir           string
	rootDir           string
	archiveDir        string
	versionDir        string
	archiveNamePrefix string
}

func NewManager() (*Manager, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	return &Manager{
		homeDir:           home,
		rootDir:           config.GoRootDirName,
		archiveDir:        config.GoArchiveDirName,
		versionDir:        config.GoVersionsDirName,
		archiveNamePrefix: config.GoVersionArchiveNamePrefix,
	}, nil
}

// ArchivePath returns the directory where archives are stored
func (m *Manager) ArchivePath() string {
	return path.Join(m.homeDir, m.rootDir, m.archiveDir)
}

// ArchiveFilePath returns the full path to a specific version's archive
func (m *Manager) ArchiveFilePath(version string) string {
	fileName := fmt.Sprintf("%s%s.tar.gz", m.archiveNamePrefix, version)
	return path.Join(m.ArchivePath(), fileName)
}

// VersionPath returns the directory where a version is installed
func (m *Manager) VersionPath(version string) string {
	return path.Join(m.homeDir, m.rootDir, m.versionDir, version)
}

func (m *Manager) DownloadVersion(version string) error {
	_, err := downloader.GetGoVersion(version, m.ArchivePath(), m.archiveNamePrefix)
	return err
}

func (m *Manager) ExtractVersion(version string) error {
	// Create the go version directory
	goVersionPath := m.VersionPath(version)
	if err := os.MkdirAll(goVersionPath, 0o755); err != nil {
		return fmt.Errorf("failed to create the version directory: %w", err)
	}
	// Decompress archived version of GO and save the content to the versioned directory
	archiveFilePath := m.ArchiveFilePath(version)
	err := decompressor.DecompressSave(archiveFilePath, goVersionPath)
	return err
}
