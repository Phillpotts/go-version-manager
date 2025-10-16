package main

import (
	"fmt"
	"os"

	"github.com/phillpotts/go-version-manager/internal/downloader"
)

func main() {
	goVersion := "2.25.3"
	path, err := downloader.DownloadGoVersion(goVersion)
	if err != nil {
		fmt.Printf("failed to download go version: %s", goVersion)
		os.Exit(1)
	}
	fmt.Println(path)
}
