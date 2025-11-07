// Package argparser provides cli command capabilities
package argparser

import (
	"errors"
	"fmt"
	"os"
)

type ArgParser struct {
	RootCommand Command
	Version     string
}

type Command struct {
	Name        string
	Description string
	Commands    []Command
	Invoke      func(args []string) error
}

func NewArgParser(name string, description string, version string) *ArgParser {
	return &ArgParser{
		Version: version,
	}
}

func newCommand(name string, description string, fn func([]string) error) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Invoke:      fn,
	}
}

func (a *ArgParser) AddCommand(name string, description string, fn func([]string) error) *Command {
	command := newCommand(name, description, fn)
	a.RootCommand.Commands = append(a.RootCommand.Commands, *command)
	return command
}

func (a *ArgParser) Parse() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("failed expected more than one argument")
	}
	return parse(args[1:], a.RootCommand)
}

func parse(args []string, command Command) error {
	if len(command.Commands) == 0 {
		return command.Invoke(args)
	}
	command, err := findCommand(args[0], command.Commands)
	if err != nil {
		return err
	}
	return parse(args[1:], command)
}

func findCommand(name string, commands []Command) (Command, error) {
	for _, command := range commands {
		if name == command.Name {
			return command, nil
		}
	}
	return Command{}, errors.New("command not found")
}
