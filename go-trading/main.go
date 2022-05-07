package main

import (
	"fmt"
	bitflyer "go-trading/bitflyer"
	config "go-trading/config"
	utils "go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
