package main

import (
	"fmt"

	"github.com/craigastuckey/gator/internal/config"
)

func main() {
	cfg := config.Read()

	cfg.SetUser("craig")

	newConfig := config.Read()

	fmt.Println(newConfig.Db_url)
	fmt.Println(newConfig.Current_user_name)

}
