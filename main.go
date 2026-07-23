package main

import (
	"fmt"

	"github.com/craigastuckey/gator/internal/config"
)

type state struct {
	conf *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		err := fmt.Errorf("login requires a single argument: login <usersname>")
		return err
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

	cfg.SetUser("craig")

	newConfig := config.Read()

	fmt.Println(newConfig.Db_url)
	fmt.Println(newConfig.Current_user_name)

}
