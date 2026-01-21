package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmd_func, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command %s\n", cmd.name)
	}
	return cmd_func(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.cmds[name] = f
	return nil
}
