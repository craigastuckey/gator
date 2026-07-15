package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string
	Current_user_name string
}

const configFileName = "gatorconfig.json"

func (c *Config) SetUser(username string) error {
	if username == "" {
		return errors.New("must provide a username")
	}
	c.Current_user_name = username

	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("error finding home directory")
	}

	path := homeDir + "/" + configFileName

	return path, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	perms := os.FileMode(0600)
	err = os.WriteFile(path, data, perms)
	if err != nil {
		return err
	}

	return nil
}

func Read() Config {
	path, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error unmarshaling config file")
		panic(err)
	}

	return cfg
}
