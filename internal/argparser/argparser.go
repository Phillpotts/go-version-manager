// Package argparser provides cli command capabilities
package argparser

import (
	"errors"
	"fmt"
	"os"

	"github.com/phillpotts/go-version-manager/internal/manager"
)

type ArgParser struct {
	RootCommand Command
	Service     manager.Manager
	Version     string
}

type Command struct {
	Name        string
	Description string
	Commands    []Command
	Invoke      func(service manager.Manager, args []string) error
}

func NewArgParser(name string, description string, version string, service manager.Manager) *ArgParser {
	return &ArgParser{
		Version: version,
		Service: service,
	}
}

func newCommand(name string, description string, fn func(manager.Manager, []string) error) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Invoke:      fn,
	}
}

func (a *ArgParser) AddCommand(name string, description string, fn func(manager.Manager, []string) error) *Command {
	command := newCommand(name, description, fn)
	a.RootCommand.Commands = append(a.RootCommand.Commands, *command)
	return command
}

func (a *ArgParser) Parse() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("failed expected more than one argument")
	}
	return parse(a.Service, args[1:], a.RootCommand)
}

func parse(service manager.Manager, args []string, command Command) error {
	if len(command.Commands) == 0 {
		return command.Invoke(service, args)
	}
	command, err := findCommand(args[0], command.Commands)
	if err != nil {
		return err
	}
	return parse(service, args[1:], command)
}

func findCommand(name string, commands []Command) (Command, error) {
	for _, command := range commands {
		if name == command.Name {
			return command, nil
		}
	}
	return Command{}, errors.New("command not found")
}
