package clio

import (
	"fmt"
)

type FlagType int64
type CommandHandler func(args map[string]FlagType) error

const (
	String FlagType = iota
	Boolean
)

type flag struct {
	name     string
	aliases  []byte
	flagType FlagType
	required bool
}

type positionalArg struct {
	name     string
	required bool
}

// When validating, command has the following required fields: name, handler, parent
type command struct {
	name           string
	flags          []*flag
	subCommands    []*command
	parent         *command
	handler        CommandHandler
	positionalArgs []*positionalArg
}

func newCommand(name string, parent *command, handler CommandHandler) *command {
	return &command{
		name:        name,
		parent:      parent,
		subCommands: make([]*command, 0),
		flags:       make([]*flag, 0),
		handler:     handler,
	}
}

func (c *command) AddNewSubCommand(name string, handler CommandHandler) *command {
	commandToAdd := newCommand(name, c, handler)
	c.subCommands = append(c.subCommands, commandToAdd)
	return commandToAdd
}

func (c *command) AddSubCommand(subCommand command) *command {
	c.subCommands = append(c.subCommands, &subCommand)
	return c
}

func (c *command) validateCommand() []error {
	errors := make([]error, 0)

	if c.name == "" {
		errors = append(errors, fmt.Errorf("Name missing from command: %+v", c))
	}

	if c.parent == nil {
		errors = append(errors, fmt.Errorf("Parent missing from command: %+v", c))
	}

	if c.handler == nil {
		errors = append(errors, fmt.Errorf("Handler missing from command: %+v", c))
	}

	return errors
}

func (c *command) run() error {
	return c.handler(make(map[string]FlagType))
}
