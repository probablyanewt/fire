package clio

import (
	"fmt"
)

type FlagType int64
type Flags map[string]FlagType
type Positionals []string
type CommandHandler func(flags Flags, positionals Positionals) error

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
	// TODO change to return error if subcommand exists?
	commandToAdd := newCommand(name, c, handler)
	c.subCommands = append(c.subCommands, commandToAdd)
	return commandToAdd
}

func (c *command) AddSubCommand(subCommand command) *command {
	// TODO change to return error if subcommand exists?
	subCommand.parent = c
	c.subCommands = append(c.subCommands, &subCommand)
	return c
}

func (c *command) validate() []error {
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

func (c *command) run(args []string) error {
	return c.handler(make(map[string]FlagType), make([]string, 0))
}

func (c *command) parseArgs(args []string) error {
	return nil
}

// getRootCommand will ascend the tree until a node without a parent is found
// It returns the root command
func (c *command) getRootCommand() *command {
	lastNode := c
	for {
		if lastNode.parent == nil {
			return lastNode
		}
		lastNode = lastNode.parent
	}
}

// flattenCommandTree flattenns the command tree from the current command into a slice of command references
// It returns a slive of command references
func (c *command) flattenCommandTree() []*command {
	commands := make([]*command, 0)
	commands = append(commands, c)
	for _, command := range commands {
		commands = append(commands, command.subCommands...)
	}
	return commands
}
