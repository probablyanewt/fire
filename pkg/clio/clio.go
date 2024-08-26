package clio

import (
	"fmt"
	"os"
)

type cli struct {
	command
}

func NewCli(name string) *cli {
	return &cli{command: command{name: name}}
}

func (c *cli) Validate() {
	validationErrors := make([]error, 0)
	for _, command := range c.flattenCommandTree() {
		validationErrors = append(validationErrors, command.validate()...)
	}

	for _, validationError := range validationErrors {
		fmt.Printf("%v", validationError)
	}

	numberOfErrors := len(validationErrors)
	if numberOfErrors > 0 {
		panic(fmt.Sprintf("%v validation errors found", numberOfErrors))
	}

	fmt.Print("Everything looks good")
}

func (c *cli) Run() {
	args := os.Args[:1]
	// Parse args to figure out which command to run and with what options
	c.run(args)
}
