package main

import (
	"fmt"
	"os"

	"github.com/phillpotts/go-version-manager/internal/argparser"
	"github.com/phillpotts/go-version-manager/internal/commands"
	"github.com/phillpotts/go-version-manager/internal/manager"
)

func main() {
	service, err := manager.NewManager()
	if err != nil {
		fmt.Printf("failed to initialize: %s\n", err)
		os.Exit(1)
	}
	argparser, err := buildArgparser(*service)
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

func buildArgparser(service manager.Manager) (argparser.ArgParser, error) {
	// Build argsparser
	app := argparser.NewArgParser("Go Version Manager", "A Cli to manage locally installed versions of the Go Language", "0.1.0-alpha", service)
	app.AddCommand("download", "download and archive the passed version", commands.DownloadVersion)
	app.AddCommand("extract", "extract an archive ready for use", commands.ExtractVersion)
	return *app, nil
}
