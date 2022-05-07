package main

import (
	"fmt"
	models "go-trading/app/models"
	config "go-trading/config"
	utils "go-trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
