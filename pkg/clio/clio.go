package clio

type cli struct {
	command
}

func NewCli(name string) *cli {
	return &cli{command: command{name: name}}
}

func (c *cli) Validate() error {
	return nil
}

func (c *cli) Run() {
}
