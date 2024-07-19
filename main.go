package main

import (
	"github.com/GabrielBruno7/gopportunities/config"
	"github.com/GabrielBruno7/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init();

	if err != nil {
		logger.ErrorFormated("Config initialization error %v", err)
		return
	}

	router.Initialize();
}

//gorm
//gin-gonic
//handler
//ponteiros
//go mod tidy //Reset the packs and install 