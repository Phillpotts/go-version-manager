package main

import (
	"fmt"
	"os"

	"github.com/phillpotts/go-version-manager/internal/downloader"
)

func main() {
	goVersion := "1.25.3"
	path, err := downloader.GetGoVersion(goVersion)
	if err != nil {
		fmt.Printf("failed to download go version: %s", goVersion)
		os.Exit(1)
	}
	fmt.Println(path)
}
