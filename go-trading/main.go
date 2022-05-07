package main

import (
	"go-trading/app/controllers"
	config "go-trading/config"
	utils "go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
