package main

import (
	"Go-Git/config"
	"Go-Git/core"
	"Go-Git/terminal"
	"bufio"
	"os"
)

func main() {
	configFilename := "./gogit.config"
	conf := config.SetConfig(configFilename)
	core.Service(conf)

	terminal.PrintNotice("Press Enter to exit...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
