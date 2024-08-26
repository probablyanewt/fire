package clio

import "testing"

func noop(_ Flags, _ Positionals) error {
	return nil
}

// TestParseArgsValid tests that a command returns no errors for valid args configuration
func TestParseArgsValid(t *testing.T) {
	command := newCommand("test", nil, noop)
	//TODO finish test
	println(command)
}
