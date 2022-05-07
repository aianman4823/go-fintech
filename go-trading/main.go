package main

import (
	config "go-trading/config"
	utils "go-trading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
