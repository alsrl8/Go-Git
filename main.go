package main

import (
	"Go-Git/config"
	"Go-Git/core"
)

func main() {
	configFilename := "./gogit.config"
	conf := config.SetConfig(configFilename)
	core.Service(conf)
}
