package cmd

import (
	"test/config"
	"test/rest"
)

func Serve() {
	config := config.GetConfig()
	rest.Start(config)
}
