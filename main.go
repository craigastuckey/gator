package main

import (
	"fmt"
	"os"

	"github.com/craigastuckey/gator/internal/config"
)

type state struct {
	conf *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s.conf == nil {
		err := fmt.Errorf("Error: state does not exist")
		return err
	}

	err := c.cmds[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("login requires a single argument: login <usersname>")
		return fmt.Errorf("login requires a single argument: login <usersname>")
	}

	err := s.conf.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}

func main() {
	cfg := config.Read()
	s := state{
		&cfg,
	}

	cmds := make(map[string]func(*state, command) error)
	c := commands{
		cmds,
	}

	c.register("login", handlerLogin)

	cmd := command{
		os.Args[1],
		os.Args[2:],
	}

	err := c.run(&s, cmd)
	if err != nil {
		panic(err)
	}
}
